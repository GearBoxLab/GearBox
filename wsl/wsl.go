package wsl

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"LeoOnTheEarth/GearBox/process"
)

var playbookNames = map[string]string{
	"Ubuntu":       "ubuntu",
	"Ubuntu-18.04": "ubuntu",
	"Ubuntu-20.04": "ubuntu",
	"Ubuntu-22.04": "ubuntu",
}

type WSL struct {
	distribution string
}

func New(distribution string) (*WSL, error) {
	if "" == distribution {
		return NewWithDefaultDistribution()
	}

	if isValid, err := ValidateDistribution(distribution); nil != err {
		return nil, err
	} else if false == isValid {
		return nil, fmt.Errorf("the distribution name %q is invalid", distribution)
	}

	return &WSL{distribution}, nil
}

func NewWithDefaultDistribution() (*WSL, error) {
	if distribution, err := GetDefaultDistribution(); nil != err {
		return nil, err
	} else {
		return &WSL{distribution}, nil
	}
}

func (w *WSL) NewProcess(name string, args ...string) *process.Process {
	p := process.New(name, args...)
	p.WSLArguments = []string{"wsl", "-d", w.distribution}

	return p
}

func (w *WSL) ConvertToLinuxPath(windowsPath string) (path string, err error) {
	var result []byte
	proc := w.NewProcess("wslpath", "-a", filepath.ToSlash(windowsPath))

	if result, err = proc.NewCommand().Output(); nil != err {
		return "", err
	}

	path = strings.TrimSpace(string(result))
	path = strings.ReplaceAll(path, "\x00", "")
	path = strings.ReplaceAll(path, "\r", "")

	return path, nil
}

func GetDefaultDistribution() (distribution string, err error) {
	var result []byte
	cmd := exec.Command("wsl", "--list", "--verbose")
	regex := regexp.MustCompile(`^\* (\S+)\s+(Running|Stopped)\s+\d`)

	if result, err = cmd.Output(); nil != err {
		return "", err
	}

	for _, line := range strings.Split(string(result), "\n") {
		line = strings.ReplaceAll(line, "\x00", "")
		line = strings.ReplaceAll(line, "\r", "")

		if matches := regex.FindStringSubmatch(line); len(matches) > 0 {
			return matches[1], nil
		}
	}

	return "", errors.New("unknown default distribution name")
}

func ValidateDistribution(distribution string) (isValid bool, err error) {
	var result []byte
	cmd := exec.Command("wsl", "--list")

	if result, err = cmd.Output(); nil != err {
		return false, err
	}

	for _, line := range strings.Split(string(result), "\n") {
		line = strings.ReplaceAll(line, "\x00", "")
		line = strings.ReplaceAll(line, "\r", "")

		if strings.Contains(line, distribution) {
			return true, nil
		}
	}

	return false, nil
}

func GetValidDistributions() []string {
	distributions := make([]string, 0)
	cmd := exec.Command("wsl", "--list")

	if result, err := cmd.Output(); nil == err {
		for _, line := range strings.Split(string(result), "\n") {
			line = strings.ReplaceAll(line, "\x00", "")
			line = strings.ReplaceAll(line, "\r", "")

			for distribution, _ := range playbookNames {
				if strings.Contains(line, distribution) {
					distributions = append(distributions, distribution)
				}
			}
		}
	}

	return distributions
}

func GetPlaybookName(distribution string) string {
	if playbookName, ok := playbookNames[distribution]; !ok {
		return ""
	} else {
		return playbookName
	}
}
