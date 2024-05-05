package ls

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:     "ls",
	Short:   "List the available ports to connect to",
	Long:    "List the available ports to connect to",
	Version: "1.0.0",
	Run:     list,
}

func list(cmd *cobra.Command, args []string) {
	// TODO: List available serial ports here
}
