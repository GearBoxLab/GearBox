package commands

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"LeoOnTheEarth/GearBox/configuration"
	"LeoOnTheEarth/GearBox/process"

	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
	sshTerminal "golang.org/x/crypto/ssh/terminal"
)

func All() []*console.Command {
	return []*console.Command{
		helpCommand,
		versionCommand,
		initCommand,
		installCommand,
		importHostsCommand,
	}
}

func getCurrentUsername(factory process.Factory) (string, error) {
	cmd := factory.NewProcess("whoami").NewCommand()

	if result, err := cmd.Output(); nil != err {
		return "", err
	} else {
		return strings.TrimSpace(string(result)), nil
	}
}

func readSudoPassword() (sudoPassword string, err error) {
	var result []byte

	terminal.Print("Enter SUDO password: ")

	if result, err = sshTerminal.ReadPassword(int(os.Stdin.Fd())); nil != err {
		return "", err
	} else {
		sudoPassword = string(result)
	}

	terminal.Print("\n\n")

	return sudoPassword, nil
}

func getConfigurationPath() string {
	confFilePath, err := configuration.Path()

	if nil != err {
		panic(err)
	}

	return confFilePath
}

func getConfigurationDir() string {
	confFilePath, err := configuration.Dir()

	if nil != err {
		panic(err)
	}

	return confFilePath
}

func isWsl() bool {
	var result []byte
	var err error

	cmd := exec.Command("uname", "-a")

	if result, err = cmd.Output(); nil != err {
		return false
	}

	return regexp.MustCompile(`(?i)microsoft|wsl`).Match(result)
}
