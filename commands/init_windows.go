//go:build windows

package commands

import (
	"LeoOnTheEarth/GearBox/configuration"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"path/filepath"
	"strings"
)

func updateEnvPath() (err error) {
	var key registry.Key
	var paths string
	var dir string

	if key, err = registry.OpenKey(registry.CURRENT_USER, "Environment", registry.QUERY_VALUE|registry.SET_VALUE); nil != err {
		return err
	}

	if paths, _, err = key.GetStringValue("PATH"); nil != err {
		return err
	}

	if dir, err = configuration.Dir(); nil != err {
		return err
	}

	path := filepath.Join(dir, "bin")

	if !strings.Contains(paths, path) {
		paths = fmt.Sprintf("%s;%s", strings.TrimRight(paths, ";"), path)

		if err = key.SetStringValue("PATH", paths); nil != err {
			return err
		}
	}

	if err = key.Close(); nil != err {
		return err
	}

	return nil
}
