package commands

import (
	"LeoOnTheEarth/GearBox/configuration"
	"fmt"

	"github.com/symfony-cli/console"
)

var initCommand = &console.Command{
	Name:  "init",
	Usage: fmt.Sprintf("Generate default configuration file (%s)", getConfigurationPath()),
	Action: func(c *console.Context) (err error) {
		if _, err = configuration.Load(); nil != err {
			return err
		}

		if err = updateEnvPath(); nil != err {
			return err
		}

		return nil
	},
}
