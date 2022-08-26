//go:build linux

package commands

import (
	"LeoOnTheEarth/GearBox/ansible"
	"LeoOnTheEarth/GearBox/configuration"
	"LeoOnTheEarth/GearBox/process"
	"errors"
	"fmt"
	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var supportedLinuxDistributions = []string{
	"ubuntu",
}

var installCommand = &console.Command{
	Name:  "install",
	Usage: "Install packages with Ansible script.",
	Action: func(c *console.Context) (err error) {
		if isWsl() {
			terminal.Print(terminal.FormatBlockMessage("error", `You must use the Windows version to run "install" command.`))
			return nil
		}

		var playbookName string
		var conf *configuration.Configuration
		var username string
		var sudoPassword string
		processFactory := process.NewFactory()

		if ok, distribution := isSupportedDistribution(); !ok {
			return fmt.Errorf("you are using unsupported distribution, use %s instead", strings.Join(supportedLinuxDistributions, ", "))
		} else {
			playbookName = distribution
		}

		if username, err = getCurrentUsername(processFactory); nil != err {
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

		ansibleInstaller := ansible.NewInstaller(processFactory)

		if err = ansibleInstaller.Install(playbookName, sudoPassword); nil != err {
			return err
		}

		var playbookMainFilePath string
		var configurationFilePath string
		var extraVarFilePath string

		if playbookMainFilePath, err = ansibleInstaller.PlaybookMainFilePath(playbookName); nil != err {
			return err
		}

		if configurationFilePath, err = configuration.Path(); nil != err {
			return err
		}

		if extraVarFilePath, err = configuration.GenerateExtraVarsFile(username, playbookName, isWsl()); nil != err {
			return err
		}

		showInstallPackages(conf)

		if true == terminal.AskConfirmation("Start to install?", false) {
			terminal.Print("\n")

			if err = ansibleInstaller.RunAnsiblePlaybook(playbookMainFilePath, configurationFilePath, extraVarFilePath, sudoPassword, conf); nil != err {
				return err
			}
		}

		if err = os.Remove(extraVarFilePath); nil != err {
			return err
		}

		return nil
	},
}

func isSupportedDistribution() (ok bool, distribution string) {
	var result []byte
	var err error

	cmd := exec.Command("cat", "/etc/os-release")

	if result, err = cmd.Output(); nil != err {
		return false, ""
	}

	regex := regexp.MustCompile(`ID=(.+)`)

	for _, line := range strings.Split(string(result), "\n") {
		line = strings.ReplaceAll(line, "\x00", "")
		line = strings.ReplaceAll(line, "\r", "")

		if matches := regex.FindStringSubmatch(line); len(matches) > 1 {
			distribution = strings.ToLower(strings.TrimSpace(matches[1]))

			for _, supportedDistribution := range supportedLinuxDistributions {
				if distribution == supportedDistribution {
					return true, distribution
				}
			}
		}
	}

	return false, ""
}
