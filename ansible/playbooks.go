package ansible

import (
	"LeoOnTheEarth/GearBox/configuration"
	"bytes"
	"crypto/md5"
	"embed"
	"errors"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed playbooks/*
var playbooks embed.FS

type InstallResult struct {
	CreatedFile string
	CreatedDir  string
}

type file struct {
	name string
	dir  string
}

func InstallPlaybookFiles(playbookName string, installDir string, conf *configuration.Configuration) (results []*InstallResult, err error) {
	var files []*file
	results = make([]*InstallResult, 0)

	if files, err = scanDir("playbooks/" + playbookName); nil != err {
		return results, err
	} else {
		var installed bool
		var result *InstallResult

		for _, f := range files {
			if installed, err = isPlaybookFileInstalled(f, installDir); nil != err {
				return results, err
			} else if false == installed {
				if result, err = installPlaybookFile(f, installDir); nil != err {
					return results, err
				} else if "" != result.CreatedFile || "" != result.CreatedDir {
					results = append(results, result)
				}
			}
		}

		if result, err = installPlaybookMainFile(playbookName, installDir, conf); nil != err {
			return results, err
		} else if "" != result.CreatedFile || "" != result.CreatedDir {
			results = append(results, result)
		}
	}

	return results, nil
}

func scanDir(dir string) (files []*file, err error) {
	var entries []fs.DirEntry
	files = make([]*file, 0)

	if entries, err = playbooks.ReadDir(dir); nil != err {
		return files, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			var subFiles []*file

			if subFiles, err = scanDir(dir + "/" + entry.Name()); nil != err {
				return files, err
			}

			for _, subFile := range subFiles {
				files = append(files, subFile)
			}
		} else {
			files = append(files, &file{name: entry.Name(), dir: dir})
		}
	}

	return files, nil
}

func isPlaybookFileInstalled(f *file, installDir string) (result bool, err error) {
	var originalContent []byte
	var targetContent []byte

	if originalContent, err = playbooks.ReadFile(filepath.ToSlash(filepath.Join(f.dir, f.name))); nil != err {
		return false, err
	}
	result1 := md5.Sum(originalContent)

	targetPath := filepath.Join(installDir, f.dir, f.name)
	if _, err = os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if targetContent, err = os.ReadFile(targetPath); nil != err {
		return false, err
	}
	result2 := md5.Sum(targetContent)

	return result1 == result2, nil
}

func installPlaybookFile(f *file, installDir string) (result *InstallResult, err error) {
	var content []byte
	originalPath := filepath.ToSlash(filepath.Join(f.dir, f.name))
	targetPath := filepath.Join(installDir, f.dir, f.name)
	targetDir := filepath.Dir(targetPath)
	result = &InstallResult{"", ""}

	if _, err = os.Stat(targetDir); errors.Is(err, os.ErrNotExist) {
		if err = os.MkdirAll(targetDir, 0755); nil != err {
			return result, err
		} else {
			result.CreatedDir = targetDir
		}
	}

	if content, err = playbooks.ReadFile(originalPath); nil != err {
		return result, err
	}

	content = bytes.Replace(content, []byte{'\r', '\n'}, []byte{'\n'}, -1)

	if err = os.WriteFile(targetPath, content, 0644); nil != err {
		return result, err
	} else {
		result.CreatedFile = targetPath
	}

	return result, nil
}

func installPlaybookMainFile(playbookName, installDir string, conf *configuration.Configuration) (result *InstallResult, err error) {
	var tmpl *template.Template
	var targetFS *os.File
	targetPath := filepath.Join(installDir, "playbooks", playbookName, "main.yml")
	result = &InstallResult{"", ""}

	if tmpl, err = template.ParseFS(playbooks, "playbooks/main.yml.gotmpl"); nil != err {
		return result, err
	}

	if targetFS, err = os.Create(targetPath); nil != err {
		return result, err
	}

	if err = tmpl.Execute(targetFS, conf); nil != err {
		return result, err
	}

	result.CreatedFile = targetPath

	return result, err
}
