package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the UI.
func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, spinner.Tick)

	return tea.Batch(cmds...)
}
