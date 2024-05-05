package cmd

import (
	"tty/cmd/connect"
	"tty/cmd/ls"

	"github.com/spf13/cobra"
)

func Execute() error {
	root := &cobra.Command{
		Use:     "tty [ls | connect]",
		Short:   "TODO",
		Long:    `TODO`,
		Version: "1.0.0",
	}

	root.AddCommand(ls.Cmd, connect.Cmd)

	return root.Execute()
}
