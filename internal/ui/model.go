package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

// Model represents the state of the UI.
type Model struct {
	loader   spinner.Model
	viewport viewport.Model
	ready    bool
}

// NewModel creates an instance of the UI.
func NewModel() Model {
	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	return Model{
		loader:   l,
		viewport: viewport.Model{},
		ready:    false,
	}
}
