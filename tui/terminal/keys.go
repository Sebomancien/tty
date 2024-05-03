package terminal

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
)

type keyMap struct {
	Help         key.Binding
	Quit         key.Binding
	Previous     key.Binding
	Next         key.Binding
	Send         key.Binding
	ScrollDown   key.Binding
	ScrollUp     key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	PageDown     key.Binding
	PageUp       key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Previous, k.Next},
		{k.ScrollUp, k.ScrollDown},
		{k.PageUp, k.PageDown},
		{k.HalfPageUp, k.HalfPageDown},
		{k.Help, k.Quit},
		{k.Send},
	}
}

func (k keyMap) HistoryKeys() viewport.KeyMap {
	return viewport.KeyMap{
		Up:           keys.ScrollUp,
		Down:         keys.ScrollDown,
		PageUp:       keys.PageUp,
		PageDown:     keys.PageDown,
		HalfPageUp:   keys.HalfPageUp,
		HalfPageDown: keys.HalfPageDown,
	}
}

var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("f1"),
		key.WithHelp("f1", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc", "quit"),
	),
	Previous: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "previous"),
	),
	Next: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "next"),
	),
	Send: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↵", "send"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys("ctrl+up"),
		key.WithHelp("ctrl+↑", "scroll up"),
	),
	ScrollDown: key.NewBinding(
		key.WithKeys("ctrl+down"),
		key.WithHelp("ctrl+↓", "scroll down"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("pgdown"),
		key.WithHelp("pgdn", "page down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup"),
		key.WithHelp("pgup", "page up"),
	),
	HalfPageUp: key.NewBinding(
		key.WithKeys("ctrl+pgdown"),
		key.WithHelp("ctrl+pgdown", "½ page up"),
	),
	HalfPageDown: key.NewBinding(
		key.WithKeys("ctrl+pgup"),
		key.WithHelp("ctrl+pgup", "½ page down"),
	),
}
