package tui

import (
	"tty/terminal"
	model "tty/tui/terminal"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(terminal *terminal.Terminal) error {
	model := model.New(terminal)
	_, err := tea.NewProgram(model).Run()
	return err
}
