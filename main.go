package main

import (
	"fmt"
	"os"
	"tty/terminal"
	"tty/tui"
)

func main() {
	terminal := terminal.New()

	err := tui.Run(terminal)
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
