package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles updating the UI.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height
		m.viewport.Width = msg.Width
		m.viewport.SetContent("Welcome to the bubbletea-starter app")

		if !m.ready {
			m.ready = true
		}

	case tea.KeyMsg:
		switch msg.String() {
		// Exit bubbletea-starter.
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.loader, cmd = m.loader.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
