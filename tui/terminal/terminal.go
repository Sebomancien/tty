package terminal

import (
	"fmt"
	"strings"
	"tty/terminal"

	"github.com/charmbracelet/bubbles/help"
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
	title    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4")).Bold(true).PaddingLeft(4).SetString("One Terminal To Rule Them All")
)

type Model struct {
	ready    bool
	keys     keyMap
	terminal *terminal.Terminal
	input    textinput.Model
	logger   viewport.Model
	help     help.Model
	history  []string
	index    int
}

func New(terminal *terminal.Terminal) Model {
	return Model{
		ready:    false,
		terminal: terminal,
		history:  []string{},
		index:    0,
		keys:     keys,
		help:     help.New(),
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
			title = title.Width(msg.Width)

			m.input = textinput.New()
			m.input.Placeholder = "Enter what you want to send here"
			m.input.Focus()
			m.input.CharLimit = 156
			m.input.Width = msg.Width

			m.logger = viewport.New(msg.Width, msg.Height-5)
			m.logger.MouseWheelEnabled = true
			m.logger.KeyMap = keys.HistoryKeys()

			m.help.Width = msg.Width

			m.ready = true
		} else {
			m.input.Width = msg.Width
			m.logger.Width = msg.Width
			m.logger.Height = msg.Height - 5
			m.help.Width = msg.Width
		}
		cmds = append(cmds, viewport.Sync(m.logger))
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Send):
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
		case key.Matches(msg, keys.Previous):
			if m.index > 0 {
				m.index--
				m.input.SetValue(m.history[m.index])
			}
		case key.Matches(msg, keys.Next):
			if m.index < len(m.history) {
				m.index++
				if m.index == len(m.history) {
					m.input.SetValue("")
				} else {
					m.input.SetValue(m.history[m.index])
				}
			}
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
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

	output := title.String() + "\n"
	output += m.logger.View() + "\n"
	output += line + "\n"
	output += m.input.View() + "\n"
	output += m.help.View(m.keys)
	if !m.help.ShowAll {
		output += "\n"
	}
	return output
}
