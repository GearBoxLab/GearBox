package ansible

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"LeoOnTheEarth/GearBox/configuration"
	"LeoOnTheEarth/GearBox/process"

	"github.com/symfony-cli/terminal"
)

type Installer struct {
	processFactory process.Factory
}

func NewInstaller(processFactory process.Factory) *Installer {
	return &Installer{
		processFactory: processFactory,
	}
}

func (i *Installer) Install(playbookName, sudoPassword string) error {
	if installed, err := isAnsibleInstalled(i.processFactory); nil != err {
		return err
	} else {
		if false == installed {
			switch playbookName {
			case "ubuntu":
				scripts := []string{
					"apt update",
					"apt install software-properties-common -y",
					"add-apt-repository --yes --update ppa:ansible/ansible",
					"apt install ansible -y",
				}
				for _, script := range scripts {
					if err = i.runSudoBashCommand(sudoPassword, script); nil != err {
						return err
					}
				}
			default:
				return errors.New(fmt.Sprintf("unsupported os: %q", playbookName))
			}
		}

		if err = i.installPlaybookFiles(playbookName); nil != err {
			return err
		}
	}

	return nil
}

func (i *Installer) installPlaybookFiles(playbookName string) error {
	var dir string
	var err error
	var results []*InstallResult

	if dir, err = i.PlaybookDir(); nil != err {
		return err
	}

	if results, err = InstallPlaybookFiles(playbookName, dir); nil != err {
		return err
	}

	if len(results) > 0 {
		terminal.Println("<success>Update Ansible Playbook files:</success>")

		for _, result := range results {
			if "" != result.CreatedDir {
				terminal.Printf("  Created directory %s\n", result.CreatedDir)
			}
			if "" != result.CreatedFile {
				terminal.Printf("  Install file %s\n", result.CreatedFile)
			}
		}
		terminal.Print("\n")
	}

	return nil
}

func (i *Installer) RunAnsiblePlaybook(playbookFilePath, configurationFilePath, extraVarFilePath, sudoPassword string) (err error) {
	p := i.processFactory.NewProcess(
		"ansible-playbook",
		playbookFilePath,
		"--extra-vars", "@"+configurationFilePath,
		"--extra-vars", "ansible_become_password="+sudoPassword,
		"--extra-vars", "@"+extraVarFilePath,
	)

	if terminal.GetLogLevel() > 1 {
		verbose := "-" + strings.Repeat("v", terminal.GetLogLevel()-1)
		p.Arguments = append(p.Arguments, &process.Argument{Value: verbose})
	}

	p.SetSecretArguments(4)

	if err = p.Run(); nil != err {
		return err
	}

	return nil
}

// PlaybookMainFilePath returns ansible playbook main.yml file path. ($HOME/.gearbox/ansible/{playbookName}/main.yml)
func (i *Installer) PlaybookMainFilePath(playbookName string) (path string, err error) {
	if path, err = i.PlaybookDir(); nil != err {
		return "", err
	}

	return filepath.Join(path, "playbooks", playbookName, "main.yml"), nil
}

// PlaybookDir returns ansible playbook root directory. ($HOME/.gearbox/ansible)
func (i *Installer) PlaybookDir() (dir string, err error) {
	if dir, err = configuration.Dir(); nil != err {
		return "", err
	}

	dir = filepath.Join(dir, "ansible")

	if _, err = os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err = os.MkdirAll(dir, 0755); nil != err {
			return dir, err
		}
	}

	return dir, nil
}

func (i *Installer) runSudoBashCommand(sudoPassword, script string) error {
	p := i.processFactory.NewProcess("echo", sudoPassword, "|", "sudo", "-S", script)
	p.SetSecretArguments(0)

	if err := p.RunBash(); nil != err {
		return err
	}

	return nil
}

func isAnsibleInstalled(processFactory process.Factory) (installed bool, err error) {
	var path string
	var realPath string

	if path, err = processFactory.NewProcess("which", "ansible").Output(); nil != err && "exit status 1" != err.Error() {
		return false, err
	}
	path = strings.TrimSpace(path)

	if "" != path {
		if realPath, err = processFactory.NewProcess("ls", path).Output(); nil != err {
			return false, err
		}

		realPath = strings.TrimSpace(realPath)
		realPath = strings.ReplaceAll(realPath, "\x00", "")
		realPath = strings.ReplaceAll(realPath, "\r", "")

		return path == realPath, nil
	}

	return false, nil
}
