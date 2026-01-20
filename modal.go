package main

import (
	"github.com/charmbracelet/lipgloss"
)

type ConfirmModal struct {
	message string
}

func NewConfirmModal() ConfirmModal {
	return ConfirmModal{
		message: "Are you sure?",
	}
}

func NewConfirmModalWithMessage(msg string) ConfirmModal {
	return ConfirmModal{
		message: msg,
	}
}

func (m ConfirmModal) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("230")).
		Padding(1, 2).
		Width(40)

	content := m.message + "\n\n" +
		"Press 'y' or Enter to confirm\n" +
		"Press 'n' or Esc to cancel"

	return style.Render(content)
}