package commands

import (
	"github.com/symfony-cli/console"
)

var helpCommand = &console.Command{
	Name:  "help",
	Usage: "Display help for a command or a category of commands",
	Args: []*console.Arg{
		{Name: "command", Optional: true},
	},
	Action: console.ShowAppHelpAction,
}
