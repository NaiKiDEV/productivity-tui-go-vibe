package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TodoItem struct {
	title     string
	completed bool
}

type TodoModel struct {
	items    []TodoItem
	selected int
	Adding   bool
	input    string
}

func NewTodoModel() TodoModel {
	return TodoModel{
		items:    []TodoItem{},
		selected: 0,
	}
}

func (m TodoModel) Init() tea.Cmd {
	return nil
}

func (m TodoModel) Update(msg tea.Msg) (TodoModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Adding {
			return m.handleAddingInput(msg)
		}
		return m.handleNormalInput(msg)
	}
	return m, nil
}

func (m TodoModel) handleAddingInput(msg tea.KeyMsg) (TodoModel, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if strings.TrimSpace(m.input) != "" {
			m.items = append(m.items, TodoItem{
				title:     strings.TrimSpace(m.input),
				completed: false,
			})
		}
		m.Adding = false
		m.input = ""
	case "esc":
		m.Adding = false
		m.input = ""
	case "backspace":
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
		}
	default:
		// Only allow printable characters for input
		if len(msg.String()) == 1 && msg.String()[0] >= 32 && msg.String()[0] <= 126 {
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m TodoModel) handleNormalInput(msg tea.KeyMsg) (TodoModel, tea.Cmd) {
	switch msg.String() {
	case "n":
		m.Adding = true
		m.input = ""
	case "up", "k":
		if m.selected > 0 {
			m.selected--
		}
	case "down", "j":
		if m.selected < len(m.items)-1 {
			m.selected++
		}
	case "enter", " ":
		if len(m.items) > 0 && m.selected < len(m.items) {
			m.items[m.selected].completed = !m.items[m.selected].completed
		}
	}
	return m, nil
}

func (m TodoModel) deleteSelected() TodoModel {
	if len(m.items) > 0 && m.selected < len(m.items) {
		m.items = append(m.items[:m.selected], m.items[m.selected+1:]...)
		if m.selected >= len(m.items) && len(m.items) > 0 {
			m.selected = len(m.items) - 1
		}
		if len(m.items) == 0 {
			m.selected = 0
		}
	}
	return m
}

func (m TodoModel) View() string {
	var b strings.Builder

	if m.Adding {
		b.WriteString("\n Add new todo:\n")
		b.WriteString(fmt.Sprintf("  > %s│\n", m.input))
		b.WriteString("\n  Press Enter to add, Esc to cancel\n")
		return b.String()
	}

	b.WriteString("\n")

	if len(m.items) == 0 {
		b.WriteString("  No todos yet. Press 'n' to add one!\n")
	} else {
		for i, item := range m.items {
			var checkbox string

			if item.completed {
				checkbox = "[✓]"
			} else {
				checkbox = "[ ]"
			}

			line := fmt.Sprintf("%s %s", checkbox, item.title)
			
			// Apply styling based on selection and completion
			if i == m.selected {
				line = lipgloss.NewStyle().
					Background(lipgloss.Color("62")).
					Foreground(lipgloss.Color("230")).
					Margin(0, 1).
					Bold(true).
					Render(line)
			} else if item.completed {
				line = lipgloss.NewStyle().
					Foreground(lipgloss.Color("243")).
					Margin(0, 1).
					Strikethrough(true).
					Render(line)
			} else {
				line = lipgloss.NewStyle().
					Foreground(lipgloss.Color("15")).
					Margin(0, 1).
					Render(line)
			}
			
			b.WriteString(line + "\n")
		}
	}

	return b.String()
}