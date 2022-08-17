package commands

import (
	"os"
	"strings"

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
