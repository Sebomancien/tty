package terminal

import (
	"fmt"
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

type Terminal struct {
	Messages []Message
}

func New() *Terminal {
	terminal := Terminal{
		Messages: []Message{},
	}

	// Test data
	for i := range 2 {
		terminal.AddUplinkMessage(fmt.Sprintf("Uplink message %d", i))
		terminal.AddDownlinkMessage(fmt.Sprintf("Downlink message %d", i))
	}

	return &terminal
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
