package serial

import (
	"log"
	config "tty/config/serial"

	"go.bug.st/serial"
)

type Port struct {
	Name string
	port serial.Port
}

func New(name string) *Port {

	mode := getConfig(name)

	port, err := serial.Open(name, mode)
	if err != nil {
		log.Fatal(err)
	}

	return &Port{
		Name: name,
		port: port,
	}
}

func getConfig(name string) *serial.Mode {
	configuration, err := config.GetConfig(name)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the parity
	parity := serial.NoParity
	switch configuration.Parity {
	case "none":
		parity = serial.NoParity
	case "odd":
		parity = serial.OddParity
	case "even":
		parity = serial.EvenParity
	case "mark":
		parity = serial.MarkParity
	case "space":
		parity = serial.SpaceParity
	default:
		log.Fatal("Invalid parity in configuration file", configuration.Parity)
	}

	// Convert to stop bits
	stopbits := serial.OneStopBit
	switch configuration.StopBits {
	case "1":
		stopbits = serial.OneStopBit
	case "1.5":
		stopbits = serial.OnePointFiveStopBits
	case "2":
		stopbits = serial.TwoStopBits
	}

	return &serial.Mode{
		BaudRate: configuration.BaudRate,
		DataBits: configuration.DataBits,
		Parity:   parity,
		StopBits: stopbits,
		InitialStatusBits: &serial.ModemOutputBits{
			RTS: configuration.RTS,
			DTR: configuration.DTR,
		},
	}
}

func applyCommandMode(mode *serial.Mode) {

}
