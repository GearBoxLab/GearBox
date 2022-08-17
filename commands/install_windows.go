//go:build windows

package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"LeoOnTheEarth/GearBox/ansible"
	"LeoOnTheEarth/GearBox/configuration"
	"LeoOnTheEarth/GearBox/wsl"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
)

var installCommand = &console.Command{
	Name: "install",
	Args: console.ArgDefinition{
		&console.Arg{
			Name:        "distribution",
			Optional:    true,
			Description: fmt.Sprintf("WSL distribution name (e.g. %s).", strings.Join(wsl.GetValidDistributions(), ", ")),
			Default:     getDefaultDistribution(),
		},
	},
	Usage: "Install packages with Ansible script.",
	Action: func(c *console.Context) (err error) {
		var distribution = c.Args().Get("distribution")
		var WSL *wsl.WSL
		var conf *configuration.Configuration
		var username string
		var sudoPassword string

		if WSL, err = wsl.New(distribution); nil != err {
			return err
		}

		if username, err = getCurrentUsername(WSL); nil != err {
			return err
		} else if "root" == username {
			return errors.New("cannot use root user")
		}

		if conf, err = configuration.Load(); nil != err {
			return err
		}

		if sudoPassword, err = readSudoPassword(); nil != err {
			return err
		}

		ansibleInstaller := ansible.NewInstaller(WSL)
		playbookName := wsl.GetPlaybookName(distribution)

		if err = ansibleInstaller.Install(playbookName, sudoPassword); nil != err {
			return err
		}

		var playbookMainFilePath string
		var configurationFilePath string
		var extraVarFilePath string
		var wslExtraVarFilePath string

		if playbookMainFilePath, err = ansibleInstaller.PlaybookMainFilePath(playbookName); nil != err {
			return err
		} else if playbookMainFilePath, err = WSL.ConvertToLinuxPath(playbookMainFilePath); nil != err {
			return err
		}

		if configurationFilePath, err = configuration.Path(); nil != err {
			return err
		} else if configurationFilePath, err = WSL.ConvertToLinuxPath(configurationFilePath); nil != err {
			return err
		}

		if extraVarFilePath, err = configuration.GenerateExtraVarsFile(username, playbookName, true); nil != err {
			return err
		} else if wslExtraVarFilePath, err = WSL.ConvertToLinuxPath(extraVarFilePath); nil != err {
			return err
		}

		showInstallPackages(conf)

		if true == terminal.AskConfirmation("Start to install?", false) {
			terminal.Print("\n")

			if err = ansibleInstaller.RunAnsiblePlaybook(playbookMainFilePath, configurationFilePath, wslExtraVarFilePath, sudoPassword); nil != err {
				return err
			}
		}

		if err = os.Remove(extraVarFilePath); nil != err {
			return err
		}

		return nil
	},
}

func getDefaultDistribution() string {
	distribution, _ := wsl.GetDefaultDistribution()

	return distribution
}
