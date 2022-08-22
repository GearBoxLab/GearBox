package commands

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"

	"LeoOnTheEarth/GearBox/configuration"
)

func updateHostsFile(hostsFilePath string) (err error) {
	var conf *configuration.Configuration
	var file *os.File

	if conf, err = configuration.Load(); nil != err {
		return err
	}

	if 0 == len(conf.ImportHostsFiles) {
		return fmt.Errorf(`the "import_hosts_files" setting in %q is empty, no hosts is imported`, getConfigurationPath())
	}

	for _, importHostsFile := range conf.ImportHostsFiles {
		var importHosts []byte
		regexpStart := regexp.MustCompile(fmt.Sprintf(`^##>>> INSERTED BY GEARBOX ## %s ## START >>>##`, importHostsFile.Name))
		regexpEnd := regexp.MustCompile(fmt.Sprintf(`^##<<< INSERTED BY GEARBOX ## %s ## END   <<<##`, importHostsFile.Name))
		hasStart := false
		hasEnd := false

		if _, err = os.Stat(importHostsFile.Path); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("the hosts file %q is not exists, check your settings in %q", importHostsFile.Path, getConfigurationPath())
		}

		if file, err = os.Open(hostsFilePath); nil != err {
			return err
		}
		buff := new(bytes.Buffer)
		hostsFileScanner := bufio.NewScanner(file)
		hostsFileScanner.Split(bufio.ScanLines)

		for hostsFileScanner.Scan() {
			line := hostsFileScanner.Text()

			if regexpStart.MatchString(line) {
				hasStart = true
				continue
			}
			if regexpEnd.MatchString(line) {
				hasEnd = true
				continue
			}

			if (false == hasStart && false == hasEnd) || (true == hasStart && true == hasEnd) {
				buff.WriteString(line)
				buff.WriteRune('\n')
			}
		}

		if err = file.Close(); nil != err {
			return err
		}

		if importHosts, err = os.ReadFile(importHostsFile.Path); nil != err {
			return err
		}
		importHosts = bytes.TrimSpace(importHosts)

		if len(importHosts) > 0 {
			buff.WriteString(fmt.Sprintf("##>>> INSERTED BY GEARBOX ## %s ## START >>>##\n", importHostsFile.Name))
			buff.Write(importHosts)
			buff.WriteRune('\n')
			buff.WriteString(fmt.Sprintf("##<<< INSERTED BY GEARBOX ## %s ## END   <<<##\n", importHostsFile.Name))
		}

		if err = os.WriteFile(hostsFilePath, buff.Bytes(), 0644); nil != err {
			return err
		}
	}

	return nil
}
