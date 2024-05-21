package connect

import (
	"fmt"
	"os"
	"tty/terminal"
	"tty/terminal/serial"
	tui "tty/tui/terminal"

	"github.com/spf13/cobra"
)

type Parity int

var Cmd = &cobra.Command{
	Use:     "connect -n NAME [-b baudrate -p parity]",
	Short:   "TODO",
	Long:    `TODO`,
	Version: "1.0.0",
	Run:     connect,
}

var (
	name     string
	baudrate uint32
	parity   string
)

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the device to connect")
	Cmd.MarkFlagRequired("name")
	Cmd.Flags().Uint32VarP(&baudrate, "baudrate", "b", 9600, "Baudrate to use for commnunicating. Default is 9600.")
	Cmd.Flags().StringVarP(&parity, "parity", "p", "none", `Parity to use for communication. Default is none. Possible values are "none", "even", "odd".`)
}

func connect(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}

	baudrate, err := cmd.Flags().GetUint32("baudrate")
	if err != nil {
		panic(err)
	}

	parity, err := cmd.Flags().GetString("parity")
	if err != nil {
		panic(err)
	}

	fmt.Println("Name     ", name)
	fmt.Println("Baudrate ", baudrate)
	fmt.Println("Parity   ", parity)

	port := serial.New(name)

	terminal := terminal.New(name, port)

	err = tui.Run(terminal)
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func getSerialParam(cmd *cobra.Command) {
	// Default parameters

	// Config file parameters

	// Console parameters
}

func newSerialParam() {

}

func getConfigSerialParam() {

}
