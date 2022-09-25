//go:build windows

package commands

import (
	"fmt"
	"time"

	"LeoOnTheEarth/GearBox/uac"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
)

var importHostsCommand = &console.Command{
	Name:  "import-hosts",
	Usage: fmt.Sprintf(`Import hosts with "import_hosts_files" setting in "%s" file. (need root privileges)`, getConfigurationPath()),
	Action: func(c *console.Context) error {
		messageFilePath := fmt.Sprintf(`%s\tmp-import-hosts-results.txt`, getConfigurationDir())

		err := uac.Prompt(messageFilePath, 10*time.Second, func() error {
			err := updateHostsFile(`C:\Windows\System32\drivers\etc\hosts`)

			if nil == err {
				terminal.Print(terminal.FormatBlockMessage("success", `Import hosts successfully.`))
				time.Sleep(200 * time.Millisecond)
			}

			return err
		})

		if nil == err {
			terminal.Print(terminal.FormatBlockMessage("success", `Import hosts successfully.`))
		}

		return err
	},
}
