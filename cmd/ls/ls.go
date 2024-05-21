package ls

import (
	"fmt"
	"os"
	"tty/tui/ports"
	"tty/tui/window"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "ls",
	Short:   "List the available ports to connect to",
	Long:    "List the available ports to connect to",
	Version: "1.0.0",
	Run:     list,
}

func list(cmd *cobra.Command, args []string) {
	ports := ports.New()
	window := window.New(ports).SetName("TTY")
	program := tea.NewProgram(window)

	_, err := program.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
