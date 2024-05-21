package window

import (
	"tty/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	background = lipgloss.Color("#7D56F4")
	foreground = lipgloss.Color("#FAFAFA")
)

var title = lipgloss.NewStyle().Foreground(foreground).Background(background).Bold(true).PaddingLeft(1)

type Window struct {
	name   string
	model  tea.Model
	width  int
	height int
	init   bool
}

func New(model tea.Model) *Window {
	return &Window{
		name:   "Title",
		model:  model,
		width:  0,
		height: 0,
		init:   false,
	}
}

func (window *Window) SetName(name string) *Window {
	window.name = name
	return window
}

func (window *Window) SetModel(model tea.Model) *Window {
	window.model = model

	// If already initialized, update the dimension and redraw
	if window.init {
		window.model.Init()
		window.updateWindowSize(window.width, window.height)
	}

	return window
}

func (window Window) Init() tea.Cmd {
	return window.model.Init()
}

func (window Window) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		window.init = true
		cmd = window.updateWindowSize(msg.Width, msg.Height)
	default:
		window.model, cmd = window.model.Update(msg)
	}

	return window, cmd
}

func (window Window) View() string {
	if !window.init {
		return "Initializing the window..."
	}

	builder := utils.NewStringBuilder()
	builder.AppendLine(title.SetString(window.name).String())
	builder.AppendLine(window.model.View())
	return builder.String()
}

func (window *Window) updateWindowSize(width int, height int) tea.Cmd {
	window.width = width
	window.height = height
	title = title.Width(width)

	// Adjust the size of the embedded window
	var cmd tea.Cmd = nil
	window.model, cmd = window.model.Update(tea.WindowSizeMsg{
		Width:  width,
		Height: height - 1,
	})

	return cmd
}
