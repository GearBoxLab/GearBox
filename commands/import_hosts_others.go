//go:build !windows

package commands

import (
	"fmt"

	"LeoOnTheEarth/GearBox/process"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
)

var importHostsCommand = &console.Command{
	Name:  "import-hosts",
	Usage: fmt.Sprintf(`Import hosts with "import_hosts_files" setting in "%s" file. (need root privileges)`, getConfigurationPath()),
	Action: func(c *console.Context) error {
		if isWsl() {
			terminal.Print(terminal.FormatBlockMessage("error", `You must use the Windows version to run "import-hosts" command.`))
			return nil
		}

		if username, err := getCurrentUsername(process.NewFactory()); nil != err {
			return err
		} else if "root" != username {
			terminal.Print(terminal.FormatBlockMessage("error", `Must be run as root. (try "sudo gearbox import-hosts")`))
			return nil
		}

		return updateHostsFile("/etc/hosts")
	},
}
