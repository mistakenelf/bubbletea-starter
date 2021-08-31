package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View returns a string representation of the entire application UI.
func (m Model) View() string {
	if !m.ready {
		return fmt.Sprintf("%s%s", m.loader.View(), "loading...")
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true).
		Italic(true).
		Render(m.viewport.View())
}
