package ports

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

var selectKey = key.NewBinding(
	key.WithKeys("enter"),
	key.WithHelp("↵", "select"),
)

var listKeys = list.KeyMap{
	ShowFullHelp: key.NewBinding(
		key.WithKeys("f1"),
		key.WithHelp("f1", "help"),
	),
	CloseFullHelp: key.NewBinding(
		key.WithKeys("f1"),
		key.WithHelp("f1", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "quit"),
	),
	CursorUp: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	CursorDown: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "prev page"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "next page"),
	),
	GoToStart: key.NewBinding(
		key.WithKeys("pgup"),
		key.WithHelp("pgup", "go to start"),
	),
	GoToEnd: key.NewBinding(
		key.WithKeys("pgdown"),
		key.WithHelp("pgdn", "go to end"),
	),
}
