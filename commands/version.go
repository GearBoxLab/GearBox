package commands

import (
	"github.com/symfony-cli/console"
)

var versionCommand = &console.Command{
	Name:  "version",
	Usage: "Display the application version",
	Action: func(c *console.Context) error {
		console.ShowVersion(c)
		return nil
	},
}
