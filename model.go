package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tab int

const (
	todoTab tab = iota
	timerTab
)

type Model struct {
	activeTab       tab
	todoModel       TodoModel
	timerModel      TimerModel
	confirmModal    ConfirmModal
	showConfirmModal bool
	width           int
	height          int
}

func NewModel() Model {
	// Try to load saved data
	todoModel, timerModel, err := loadAppData()
	if err != nil {
		// If loading fails, use defaults
		todoModel = NewTodoModel()
		timerModel = NewTimerModel()
	}
	
	return Model{
		activeTab:    todoTab,
		todoModel:    todoModel,
		timerModel:   timerModel,
		confirmModal: NewConfirmModal(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.todoModel.Init(),
		m.timerModel.Init(),
		tea.EnterAltScreen,
		// Start the global tick for timers
		tea.Every(time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		}),
		// Auto-save every 30 seconds
		tea.Every(30*time.Second, func(t time.Time) tea.Msg {
			return AutoSaveMsg(t)
		}),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		
		var cmd tea.Cmd
		m.todoModel, cmd = m.todoModel.Update(msg)
		cmds = append(cmds, cmd)
		
		m.timerModel, cmd = m.timerModel.Update(msg)
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		if m.showConfirmModal {
			return m.handleConfirmModal(msg)
		}

		// Check if we're in adding mode - if so, skip global navigation
		inAddingMode := (m.activeTab == todoTab && m.todoModel.Adding) || 
						(m.activeTab == timerTab && m.timerModel.Adding)
		
		if !inAddingMode {
			switch msg.String() {
			case "ctrl+c", "q":
				// Save data before quitting
				saveAppData(m.todoModel, m.timerModel)
				return m, tea.Quit
			
			case "tab", "h", "l", "right", "left":
				m = m.switchTab(msg.String())
				return m, nil
			
			case "d":
				return m.handleDelete()
			}
		} else {
			// In adding mode, only allow global quit commands
			switch msg.String() {
			case "ctrl+c", "q":
				// Save data before quitting
				saveAppData(m.todoModel, m.timerModel)
				return m, tea.Quit
			}
		}
		
		return m.handleTabInput(msg)
	
	case TickMsg:
		// Always update timers regardless of active tab
		var cmd tea.Cmd
		m.timerModel, cmd = m.timerModel.Update(msg)
		cmds = append(cmds, cmd)
		
		// Regenerate the tick for continuous updates
		cmds = append(cmds, tea.Every(time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		}))
		
	case AutoSaveMsg:
		// Auto-save data periodically
		saveAppData(m.todoModel, m.timerModel)
		
		// Regenerate auto-save
		cmds = append(cmds, tea.Every(30*time.Second, func(t time.Time) tea.Msg {
			return AutoSaveMsg(t)
		}))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) switchTab(key string) Model {
	switch key {
	case "tab", "right", "l":
		if m.activeTab == todoTab {
			m.activeTab = timerTab
		} else {
			m.activeTab = todoTab
		}
	case "left", "h":
		if m.activeTab == timerTab {
			m.activeTab = todoTab
		} else {
			m.activeTab = timerTab
		}
	}
	return m
}

func (m Model) handleDelete() (tea.Model, tea.Cmd) {
	var hasItems bool
	var itemName string

	if m.activeTab == todoTab {
		hasItems = len(m.todoModel.items) > 0 && m.todoModel.selected < len(m.todoModel.items)
		if hasItems {
			itemName = m.todoModel.items[m.todoModel.selected].title
		}
	} else {
		hasItems = len(m.timerModel.items) > 0 && m.timerModel.selected < len(m.timerModel.items)
		if hasItems {
			itemName = m.timerModel.items[m.timerModel.selected].name
		}
	}

	if hasItems {
		m.confirmModal = NewConfirmModalWithMessage("Delete \"" + itemName + "\"?")
		m.showConfirmModal = true
	}

	return m, nil
}

func (m Model) handleConfirmModal(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "enter":
		m.showConfirmModal = false
		if m.activeTab == todoTab {
			m.todoModel = m.todoModel.deleteSelected()
		} else {
			m.timerModel = m.timerModel.deleteSelected()
		}
	case "n", "esc":
		m.showConfirmModal = false
	}
	return m, nil
}

func (m Model) handleTabInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	if m.activeTab == todoTab {
		m.todoModel, cmd = m.todoModel.Update(msg)
	} else {
		m.timerModel, cmd = m.timerModel.Update(msg)
	}
	
	return m, cmd
}

func (m Model) View() string {
	if m.showConfirmModal {
		return m.renderWithModal()
	}

	return m.renderMain()
}

func (m Model) renderMain() string {
	// Render title
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("230")).
		Bold(true).
		Padding(1, 1, 0, 1).
		Render("Productivity TUI")
	
	// Render tabs
	tabBar := m.renderTabBar()
	
	// Render active tab content
	var content string
	if m.activeTab == todoTab {
		content = m.todoModel.View()
	} else {
		content = m.timerModel.View()
	}
	
	// Render help at bottom
	help := m.renderHelp()

	return title + "\n\n" + tabBar + "\n\n" + content + "\n" + help
}

func (m Model) renderWithModal() string {
	base := m.renderMain()
	modal := m.confirmModal.View()
	
	// Simple overlay - place modal at the bottom
	return base + "\n\n" + modal
}

func (m Model) renderTabBar() string {
	todoStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Margin(0, 1)
	
	timerStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Margin(0, 1)

	if m.activeTab == todoTab {
		todoStyle = todoStyle.
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Bold(true)
	} else {
		todoStyle = todoStyle.
			Background(lipgloss.Color("240")).
			Foreground(lipgloss.Color("250"))
	}

	if m.activeTab == timerTab {
		timerStyle = timerStyle.
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Bold(true)
	} else {
		timerStyle = timerStyle.
			Background(lipgloss.Color("240")).
			Foreground(lipgloss.Color("250"))
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		todoStyle.Render("Todo"),
		timerStyle.Render("Timer"),
	)
}

func (m Model) renderHelp() string {
	if m.showConfirmModal {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			Render("Press y/enter to confirm • n/esc to cancel")
	}
	
	if m.activeTab == todoTab && m.todoModel.Adding {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			Render("Type todo title • enter to save • esc to cancel")
	}
	
	if m.activeTab == timerTab && m.timerModel.Adding {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			Render("Type timer name • enter to save • esc to cancel")
	}
	
	keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Bold(true)
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
	
	if m.activeTab == todoTab {
		return " " + keyStyle.Render("tab/h/l") + textStyle.Render(" switch tabs • ") + 
			keyStyle.Render("n") + textStyle.Render(" add • ") + 
			keyStyle.Render("d") + textStyle.Render(" delete • ") + 
			keyStyle.Render("j/k/↑/↓") + textStyle.Render(" navigate • ") + 
			keyStyle.Render("space/enter") + textStyle.Render(" toggle • ") + 
			keyStyle.Render("q") + textStyle.Render(" quit")
	} else {
		return " " + keyStyle.Render("tab/h/l") + textStyle.Render(" switch tabs • ") + 
			keyStyle.Render("n") + textStyle.Render(" add • ") + 
			keyStyle.Render("d") + textStyle.Render(" delete • ") + 
			keyStyle.Render("j/k/↑/↓") + textStyle.Render(" navigate • ") + 
			keyStyle.Render("space/enter") + textStyle.Render(" start/stop • ") + 
			keyStyle.Render("r") + textStyle.Render(" reset • ") + 
			keyStyle.Render("q") + textStyle.Render(" quit")
	}
}