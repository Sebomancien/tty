package terminal

import (
	"time"
)

type MessageDirection int

const (
	Up MessageDirection = iota
	Down
)

type Message struct {
	Direction MessageDirection
	Message   string
	Timestamp time.Time
}

type Port interface{}

type Terminal struct {
	Name     string
	Port     Port
	Messages []Message
}

func New(name string, port Port) *Terminal {
	return &Terminal{
		Name:     name,
		Port:     port,
		Messages: []Message{},
	}
}

func (t *Terminal) Clear() {
	t.Messages = []Message{}
}

func (t *Terminal) AddUplinkMessage(message string) {
	t.Messages = append(t.Messages, Message{
		Direction: Up,
		Message:   message,
		Timestamp: time.Now(),
	})
}

func (t *Terminal) AddDownlinkMessage(message string) {
	t.Messages = append(t.Messages, Message{
		Direction: Down,
		Message:   message,
		Timestamp: time.Now(),
	})
}
