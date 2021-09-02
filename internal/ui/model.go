package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// keyMap struct contains all keybindings.
type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

// keys represents the key bindings in the app along with the help text.
var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
	}
}

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
	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	h := help.NewModel()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	v := viewport.Model{}
	v.SetContent("Welcome to the bubbletea-starter app")

	return Model{
		keys:     keys,
		help:     h,
		loader:   l,
		viewport: v,
		ready:    false,
	}
}
