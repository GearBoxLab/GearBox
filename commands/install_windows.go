//go:build windows

package commands

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"LeoOnTheEarth/GearBox/ansible"
	"LeoOnTheEarth/GearBox/configuration"
	"LeoOnTheEarth/GearBox/wsl"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
)

var windowsFilePathRegexp = regexp.MustCompile(`^[a-zA-Z]:\\`)

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
	Flags: installCommandFlags,
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
		if err = convertConfigurationFilePaths(conf, WSL); nil != err {
			return err
		}
		if c.Bool("only-run-extra-ansible-playbooks") {
			conf.OnlyRunExtraAnsiblePlaybooks = true
		}

		sudoPassword = c.String("sudo-password")
		if 0 == len(sudoPassword) {
			if sudoPassword, err = readSudoPassword(); nil != err {
				return err
			}
		}

		ansibleInstaller := ansible.NewInstaller(WSL)
		playbookName := wsl.GetPlaybookName(distribution)

		if err = ansibleInstaller.Install(playbookName, sudoPassword, conf); nil != err {
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

		if !c.Bool("only-run-extra-ansible-playbooks") {
			showInstallPackages(conf)
		}

		if c.Bool("yes") || true == terminal.AskConfirmation("Start to install?", false) {
			terminal.Print("\n")

			if err = ansibleInstaller.RunAnsiblePlaybook(playbookMainFilePath, configurationFilePath, wslExtraVarFilePath, sudoPassword, conf); nil != err {
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

func convertConfigurationFilePaths(conf *configuration.Configuration, WSL *wsl.WSL) (err error) {
	var newPath string

	for i, path := range conf.ExtraAnsiblePlaybooks.PlaybookFiles {
		if windowsFilePathRegexp.MatchString(path) {
			if newPath, err = WSL.ConvertToLinuxPath(path); nil != err {
				return err
			}
			conf.ExtraAnsiblePlaybooks.PlaybookFiles[i] = newPath
		}
	}

	for i, path := range conf.ExtraAnsiblePlaybooks.VariableFiles {
		if windowsFilePathRegexp.MatchString(path) {
			if newPath, err = WSL.ConvertToLinuxPath(path); nil != err {
				return err
			}
			conf.ExtraAnsiblePlaybooks.VariableFiles[i] = newPath
		}
	}

	return nil
}
