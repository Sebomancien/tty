package terminal

import (
	"fmt"
	"strings"
	"tty/terminal"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	maxHistoryLen = 100
)

var (
	uplink   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00d787")).SetString("↑")
	downlink = lipgloss.NewStyle().Foreground(lipgloss.Color("#0087d7")).SetString("↓")
)

var historyKeys = viewport.KeyMap{
	PageDown: key.NewBinding(
		key.WithKeys("pgdown"),
		key.WithHelp("pgdn", "page down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup"),
		key.WithHelp("pgup", "page up"),
	),
	HalfPageUp: key.NewBinding(
		key.WithKeys("ctrl+shift+down"),
		key.WithHelp("ctrl+shift+↓", "½ page up"),
	),
	HalfPageDown: key.NewBinding(
		key.WithKeys("ctrl+shift+up"),
		key.WithHelp("ctrl+shift+↑", "½ page down"),
	),
	Up: key.NewBinding(
		key.WithKeys("ctrl+up"),
		key.WithHelp("ctrl+↑", "scroll up"),
	),
	Down: key.NewBinding(
		key.WithKeys("ctrl+down"),
		key.WithHelp("ctrl+↓", "scroll down"),
	),
}

type Model struct {
	ready    bool
	terminal *terminal.Terminal
	input    textinput.Model
	logger   viewport.Model
	history  []string
	index    int
}

func New(terminal *terminal.Terminal) Model {
	return Model{
		ready:    false,
		terminal: terminal,
		history:  []string{},
		index:    0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.input = textinput.New()
			m.input.Placeholder = "Enter what you want to send here"
			m.input.Focus()
			m.input.CharLimit = 156
			m.input.Width = msg.Width

			m.logger = viewport.New(msg.Width, msg.Height-2)
			m.logger.MouseWheelEnabled = true
			m.logger.KeyMap = historyKeys

			m.ready = true
		} else {
			m.input.Width = msg.Width
			m.logger.Width = msg.Width
			m.logger.Height = msg.Height - 2
		}
		cmds = append(cmds, viewport.Sync(m.logger))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			text := m.input.Value()
			if len(text) > 0 {
				m.terminal.AddDownlinkMessage(text)
				m.input.SetValue("")
				m.history = append(m.history, text)
				if len(m.history) > maxHistoryLen {
					m.history = m.history[:maxHistoryLen]
				}
				m.index = len(m.history)
			}
		case "up":
			if m.index > 0 {
				m.index--
				m.input.SetValue(m.history[m.index])
			}
		case "down":
			if m.index < len(m.history) {
				m.index++
				if m.index == len(m.history) {
					m.input.SetValue("")
				} else {
					m.input.SetValue(m.history[m.index])
				}
			}
		}
	}

	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)
	m.logger, cmd = m.logger.Update(msg)
	cmds = append(cmds, cmd)

	if len(cmds) > 0 {
		return m, tea.Batch(cmds...)
	} else {
		return m, nil
	}
}

func (m Model) View() string {
	if !m.ready {
		return "Initializing terminal..."
	}

	// Format each message line
	var s string
	for _, msg := range m.terminal.Messages {
		var icon lipgloss.Style
		switch msg.Direction {
		case terminal.Up:
			icon = uplink
		case terminal.Down:
			icon = downlink
		}
		s += fmt.Sprintf("%s %s: %s\n", icon, msg.Timestamp.Format("15:04:05.000"), msg.Message)
	}

	m.logger.SetContent(s)
	m.logger.GotoBottom()
	line := strings.Repeat("─", m.logger.Width)

	return fmt.Sprintf("%s\n%s\n%s", m.logger.View(), line, m.input.View())
}
