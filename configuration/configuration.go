package configuration

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const DirectoryName = ".gearbox"
const FileName = "config.json"
const ExtraVarsFileName = "extra-vars.json"

type Configuration struct {
	PHP                          confPHP                   `json:"php"`
	Blackfire                    confBlackfire             `json:"blackfire"`
	NodeJS                       confNodeJS                `json:"nodejs"`
	GoLang                       confGoLang                `json:"golang"`
	Nginx                        confNginx                 `json:"nginx"`
	Memcached                    confMemcached             `json:"memcached"`
	Redis                        confRedis                 `json:"redis"`
	ImportHostsFiles             []confImportHostFile      `json:"import_hosts_files"`
	ExtraAnsiblePlaybooks        confExtraAnsiblePlaybooks `json:"extra_ansible_playbooks"`
	ExtraServiceNames            []string                  `json:"extra_service_names"`
	onlyRunExtraAnsiblePlaybooks bool
}

type confPHP struct {
	Install        bool     `json:"install"`
	Versions       []string `json:"versions"`
	DefaultVersion string   `json:"default_version"`
	EnableService  bool     `json:"enable_service"`
}

type confBlackfire struct {
	Install       bool   `json:"install"`
	Collector     string `json:"collector"`
	LogLevel      int    `json:"log_level"`
	ServerId      string `json:"server_id"`
	ServerToken   string `json:"server_token"`
	Socket        string `json:"socket"`
	EnableService bool   `json:"enable_service"`
}

type confNodeJS struct {
	Install    bool   `json:"install"`
	Version    string `json:"version"`
	NvmVersion string `json:"nvm_version"`
}

type confGoLang struct {
	Install bool   `json:"install"`
	Version string `json:"version"`
}

type confNginx struct {
	Install        bool   `json:"install"`
	CompileVersion string `json:"compile_version"`
	EnableService  bool   `json:"enable_service"`
}

type confMemcached struct {
	Install       bool `json:"install"`
	EnableService bool `json:"enable_service"`
}

type confRedis struct {
	Install       bool `json:"install"`
	EnableService bool `json:"enable_service"`
}

type confImportHostFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type confExtraAnsiblePlaybooks struct {
	PlaybookFiles []string `json:"playbook_files"`
	VariableFiles []string `json:"variable_files"`
}

func (c *Configuration) ToJson() string {
	if content, err := json.Marshal(c); nil != err {
		return ""
	} else {
		return string(content)
	}
}

// New returns a Configuration instance with default values.
func New() *Configuration {
	return &Configuration{
		PHP: confPHP{
			Install:        true,
			Versions:       []string{"8.2"},
			DefaultVersion: "8.2",
			EnableService:  true,
		},
		Blackfire: confBlackfire{
			Install:       false,
			Collector:     "https://blackfire.io",
			LogLevel:      1,
			ServerId:      "",
			ServerToken:   "",
			Socket:        "unix:///var/run/blackfire/agent.sock",
			EnableService: true,
		},
		NodeJS: confNodeJS{
			Install:    true,
			Version:    "lts",
			NvmVersion: "0.39.3",
		},
		GoLang: confGoLang{
			Install: false,
			Version: "1.19.7",
		},
		Nginx: confNginx{
			Install:        true,
			CompileVersion: "",
			EnableService:  true,
		},
		Memcached: confMemcached{
			Install:       true,
			EnableService: true,
		},
		Redis: confRedis{
			Install:       true,
			EnableService: true,
		},
		ImportHostsFiles: []confImportHostFile{},
		ExtraAnsiblePlaybooks: confExtraAnsiblePlaybooks{
			PlaybookFiles: []string{},
			VariableFiles: []string{},
		},
		ExtraServiceNames:            []string{},
		onlyRunExtraAnsiblePlaybooks: false,
	}
}

func (c *Configuration) SetOnlyRunExtraAnsiblePlaybooks(value bool) {
	c.onlyRunExtraAnsiblePlaybooks = value
}

func (c *Configuration) OnlyRunExtraAnsiblePlaybooks() bool {
	return c.onlyRunExtraAnsiblePlaybooks
}

func (c *Configuration) Normalize() {
	phpVersions := make([]string, 0)
	for _, version := range c.PHP.Versions {
		if version != c.PHP.DefaultVersion {
			phpVersions = append(phpVersions, version)
		}
	}
	phpVersions = append(phpVersions, c.PHP.DefaultVersion)
	c.PHP.Versions = phpVersions
}

// Load loads configuration from config.json file.
func Load() (c *Configuration, err error) {
	var content []byte
	var path string

	if path, err = Path(); nil != err {
		return nil, err
	}

	c = New()

	if _, err = os.Stat(path); false == errors.Is(err, os.ErrNotExist) {
		if content, err = os.ReadFile(path); nil != err {
			return nil, err
		}
		if err = json.Unmarshal(content, c); nil != err {
			return nil, err
		}
	}

	c.Normalize()

	if content, err = json.MarshalIndent(c, "", "  "); nil != err {
		return nil, err
	}
	if err = os.WriteFile(path, content, 0644); nil != err {
		return nil, err
	}

	return c, nil
}

// Path returns configuration file path.
func Path() (path string, err error) {
	if path, err = Dir(); nil != err {
		return "", err
	}

	return filepath.Join(path, FileName), nil
}

// Dir returns configuration directory path.
func Dir() (dir string, err error) {
	if dir, err = os.UserHomeDir(); nil != err {
		return "", err
	}

	dir = filepath.Join(dir, DirectoryName)

	if _, err = os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err = os.MkdirAll(dir, 0755); nil != err {
			return "", err
		}
	}

	return dir, nil
}

func GetServiceNames(playbookName string) []string {
	names := make([]string, 0)

	if conf, _ := Load(); nil != conf {
		if "ubuntu" == playbookName {
			if conf.PHP.Install {
				for _, version := range conf.PHP.Versions {
					names = append(names, "php"+version+"-fpm")
				}
			}
			if conf.Blackfire.Install {
				names = append(names, "blackfire-agent")
			}
			if conf.Nginx.Install {
				names = append(names, "nginx")
			}
			if conf.Memcached.Install {
				names = append(names, "memcached")
			}
			if conf.Redis.Install {
				names = append(names, "redis-server")
			}
		}

		for _, name := range conf.ExtraServiceNames {
			names = append(names, name)
		}
	}

	return names
}

func GenerateExtraVarsFile(username, playbookName string, isWsl bool) (path string, err error) {
	var content []byte

	if path, err = Dir(); nil != err {
		return "", err
	}
	path = filepath.Join(path, ExtraVarsFileName)

	vars := struct {
		Username     string   `json:"username"`
		WSL          bool     `json:"wsl"`
		ServiceNames []string `json:"service_names"`
	}{
		Username:     username,
		WSL:          isWsl,
		ServiceNames: GetServiceNames(playbookName),
	}

	if content, err = json.Marshal(vars); nil != err {
		return "", err
	}

	if err = os.WriteFile(path, content, 0644); nil != err {
		return "", err
	}

	return path, nil
}
