package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the state of the UI.
type Model struct {
	keys     keyMap
	help     help.Model
	loader   spinner.Model
	viewport viewport.Model
	ready    bool
}

// NewModel creates an instance of the UI.
func NewModel() Model {
	keys := getDefaultKeyMap()

	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	h := help.NewModel()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	return Model{
		keys:     keys,
		help:     h,
		loader:   l,
		viewport: viewport.Model{},
		ready:    false,
	}
}
