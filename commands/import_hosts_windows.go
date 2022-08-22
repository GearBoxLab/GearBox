//go:build windows

package commands

import (
	"errors"
	"fmt"
	"os"
	"time"

	"LeoOnTheEarth/GearBox/uac"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
)

var importHostsCommand = &console.Command{
	Name:  "import-hosts",
	Usage: fmt.Sprintf(`Import hosts with "import_hosts_files" setting in "%s" file. (need root privileges)`, getConfigurationPath()),
	Flags: []console.Flag{
		&console.StringFlag{Name: "micro-timestamp", Required: false, Usage: "Required when running in UAC mode.", DefaultValue: fmt.Sprintf("%d", time.Now().UnixMilli())},
	},
	Action: func(c *console.Context) (err error) {
		var messageFileInfo os.FileInfo
		microTimestamp := c.String("micro-timestamp")
		messageFilePath := fmt.Sprintf(`%s\tmp-%s.txt`, getConfigurationDir(), microTimestamp)
		message := []byte("")

		if !uac.IsAdmin() {
			if err = uac.PromptWithExtraArguments([]string{"--micro-timestamp", microTimestamp}); nil != err {
				return err
			}
		} else {
			if err = updateHostsFile(`C:\Windows\System32\drivers\etc\hosts`); nil != err {
				message = []byte(err.Error())
			}

			if writeErr := os.WriteFile(messageFilePath, message, 0644); nil != writeErr {
				return writeErr
			}

			return err
		}

		spent := 0

		for true {
			time.Sleep(200 * time.Millisecond)
			spent += 200

			if messageFileInfo, err = os.Stat(messageFilePath); nil == messageFileInfo {
				if spent >= 10000 {
					break
				}
				continue
			} else {
				if message, err = os.ReadFile(messageFilePath); nil != err {
					return err
				} else if len(message) > 0 {
					err = errors.New(string(message))
				}

				if removeErr := os.Remove(messageFilePath); nil != removeErr {
					return removeErr
				}

				if nil != err {
					return err
				}
			}

			break
		}

		terminal.Print(terminal.FormatBlockMessage("success", `Import hosts successfully.`))

		return nil
	},
}
