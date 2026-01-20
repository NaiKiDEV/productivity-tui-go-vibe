package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TickMsg time.Time
type AutoSaveMsg time.Time

type TimerItem struct {
	name    string
	elapsed time.Duration
	running bool
}

type TimerModel struct {
	items    []TimerItem
	selected int
	Adding   bool
	input    string
}

func NewTimerModel() TimerModel {
	return TimerModel{
		items:    []TimerItem{},
		selected: 0,
	}
}

func (m TimerModel) Init() tea.Cmd {
	return nil
}

func (m TimerModel) Update(msg tea.Msg) (TimerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Adding {
			return m.handleAddingInput(msg)
		}
		return m.handleNormalInput(msg)
	case TickMsg:
		return m.handleTick(), nil
	}
	return m, nil
}

func (m TimerModel) handleTick() TimerModel {
	for i := range m.items {
		if m.items[i].running {
			m.items[i].elapsed += time.Second
		}
	}
	return m
}

func (m TimerModel) handleAddingInput(msg tea.KeyMsg) (TimerModel, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if strings.TrimSpace(m.input) != "" {
			m.items = append(m.items, TimerItem{
				name:    strings.TrimSpace(m.input),
				elapsed: 0,
				running: false,
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

func (m TimerModel) handleNormalInput(msg tea.KeyMsg) (TimerModel, tea.Cmd) {
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
			m.items[m.selected].running = !m.items[m.selected].running
		}
	case "r":
		// Reset timer
		if len(m.items) > 0 && m.selected < len(m.items) {
			m.items[m.selected].elapsed = 0
			m.items[m.selected].running = false
		}
	}
	return m, nil
}

func (m TimerModel) deleteSelected() TimerModel {
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

func (m TimerModel) View() string {
	var b strings.Builder

	if m.Adding {
		b.WriteString("\n Add new timer:\n")
		b.WriteString(fmt.Sprintf("  Name: %s│\n", m.input))
		b.WriteString("\n  Press Enter to add, Esc to cancel\n")
		return b.String()
	}

	if len(m.items) == 0 {
		b.WriteString(" No timers yet. Press 'n' to add one!\n")
	} else {
		for i, item := range m.items {
			// First line: name with cursor
			nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Bold(true)
			if i == m.selected {
				nameStyle = nameStyle.Background(lipgloss.Color("62")).Foreground(lipgloss.Color("230"))
			}
			nameLine := nameStyle.Render(item.name)
			
			// Second line: status and time
			minutes := int(item.elapsed.Minutes())
			seconds := int(item.elapsed.Seconds()) % 60
			timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)
			
			var statusText string
			if item.running {
				statusText = "Running"
			} else {
				statusText = "Stopped"
			}
			
			statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
			timeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Bold(true)
			if item.running {
				statusStyle = statusStyle.Foreground(lipgloss.Color("62"))
				timeStyle = timeStyle.Foreground(lipgloss.Color("62"))
			}
			
			timeLine := " " + statusStyle.Render(statusText) + " • " + timeStyle.Render(timeStr)
			
			b.WriteString(" " + nameLine + "\n")
			b.WriteString(" " + timeLine + "\n")
			b.WriteString("\n") // Extra spacing between timers
		}
	}

	return b.String()
}