package commands

import (
	"LeoOnTheEarth/GearBox/configuration"

	"github.com/symfony-cli/terminal"
)

func showInstallPackages(conf *configuration.Configuration) {
	terminal.Print("Install packages:\n")

	if conf.PHP.Install {
		for _, version := range conf.PHP.Versions {
			if conf.PHP.DefaultVersion == version {
				terminal.Printf("  - PHP v%s (default)\n", version)
			} else {
				terminal.Printf("  - PHP v%s\n", version)
			}
		}
	}

	if conf.Blackfire.Install {
		terminal.Print("  - Blackfire Agent\n")
	}

	if conf.GoLang.Install {
		terminal.Printf("  - Golang v%s\n", conf.GoLang.Version)
	}

	if conf.NodeJS.Install {
		terminal.Printf("  - NodeJS v%s\n", conf.NodeJS.Version)
	}

	if conf.Nginx.Install {
		terminal.Print("  - Nginx Server\n")
	}

	if conf.Memcached.Install {
		terminal.Print("  - Memcached Server\n")
	}

	if conf.Redis.Install {
		terminal.Print("  - Redis Server\n")
	}

	terminal.Print("\n")
}