package commands

import (
	"fmt"

	"LeoOnTheEarth/GearBox/configuration"

	"github.com/symfony-cli/console"
)

var initCommand = &console.Command{
	Name:  "init",
	Usage: fmt.Sprintf("Generate default configuration file (%s)", getConfigurationPath()),
	Action: func(c *console.Context) error {
		_, err := configuration.Load()

		return err
	},
}

func getConfigurationPath() string {
	confFilePath, _ := configuration.Path()

	return confFilePath
}
