package ports

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"go.bug.st/serial/enumerator"
)

type Ports struct {
	list list.Model
	init bool
}

func New() *Ports {
	return &Ports{
		init: false,
	}
}

func (p Ports) Init() tea.Cmd {
	return nil
}

func (p Ports) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if p.init {
			p.list.SetWidth(msg.Width)
			p.list.SetHeight(msg.Height)
		} else {
			p.list = list.New(getPorts(), list.NewDefaultDelegate(), msg.Width, msg.Height)
			p.list.SetShowTitle(false)
			p.list.KeyMap = listKeys
			p.init = true
		}
		return p, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			return p, tea.Quit

			/*
				case "enter":
					i, ok := m.list.SelectedItem().(item)
					if ok {
						m.choice = string(i)
					}
					return m, tea.Quit
			*/
		}
	}

	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p Ports) View() string {
	if !p.init {
		return "Not initialized yet"
	}

	return p.list.View()
}

func getPorts() []list.Item {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}

	var list []list.Item
	for _, port := range ports {
		list = append(list, item{title: port.Name, description: port.Product})
	}

	return list
}
