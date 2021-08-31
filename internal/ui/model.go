package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	loader   spinner.Model
	viewport viewport.Model
	ready    bool
}

func NewModel() Model {
	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	return Model{
		loader:   l,
		viewport: viewport.Model{},
		ready:    false,
	}
}
