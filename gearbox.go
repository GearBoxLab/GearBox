package main

import (
	"os"

	"LeoOnTheEarth/GearBox/commands"
	"github.com/symfony-cli/console"
)

var Version = "dev"

func main() {
	app := &console.Application{
		Name:     "GearBox",
		Usage:    "GearBox helps to easily create the same web development environment in Windows and Linux.",
		Version:  Version,
		Channel:  "stable",
		Commands: commands.All(),
	}

	app.Run(os.Args)
}
