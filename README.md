GearBox
=======

GearBox helps to easily create the same web development environment in Windows(WSL) and Linux.

## Features

1. Supported OS:
   - Windows 10/11 with WSL enabled. Supported WSL distributions:
     - Ubuntu
     - Ubuntu-18.04
     - Ubuntu-20.04
   - Linux distributions:
     - Ubuntu
2. Install packages:
   - PHP
     - Install multiple versions of PHP, v5.6 ~ v8.1
     - Provide a `phpvm` command to switch PHP versions. Usage:
       ```
       Usage: phpvm use <php-version>
       Example1: phpvm use 5.6
                 then run "php --version" will get "PHP 5.6.x"
       Example2: phpvm use 7.4
                 then run "php --version" will get "PHP 7.4.x"
       ```
   - Blackfire
   - NodeJS (include Yarn)
   - GoLang
   - Nginx
   - Memcached
   - Redis
3. Run personal Ansible Tasks when executing the `gearbox install` command.
4. Import extra hosts.
   - In Linux, will import to `/etc/hosts` file.
   - In Windows, will import to `C:\Windows\System32\drivers\etc\hosts` file.

## Install/Update GearBox in Windows

### Requirements

1. Need Windows 10 version 2004 and higher (Build 19041 and higher) or Windows 11.
2. Enable [WSL (Windows Subsystem for Linux)](https://docs.microsoft.com/en-us/windows/wsl/install-manual).
3. Set up WSL default version to 1. (It is recommended to use WSL1 for [better filesystem performance](https://docs.microsoft.com/en-us/windows/wsl/compare-versions#comparing-features))
   ```bash
   # Set WSL default version to 1.
   wsl --set-default-version 1
   ```
4. [Install a Linux distribution](https://docs.microsoft.com/en-us/windows/wsl/install).

### Installation

Run the following command in your terminal to install/update GearBox files.

```bash
bitsadmin.exe /transfer "Download GearBox Installer" https://raw.githubusercontent.com/GearBoxLab/GearBox/master/scripts/install-gearbox-windows.bat %TEMP%\install-gearbox-windows.bat && CALL %TEMP%\install-gearbox-windows.bat && DEL /Q %TEMP%\install-gearbox-windows.bat 
```

## Install/Update GearBox in Linux

Run the following command in your terminal to install/update GearBox files.

```bash
curl -s -o- https://raw.githubusercontent.com/GearBoxLab/GearBox/master/scripts/install-gearbox-linux.sh | bash
```

## Configuration

Configuration file is located in `$HOME/.gearbox/config.json` (or `%USERPROFILE%\.gearbox\config.json` in Windows)

Default configuration as below:

```json
{
  "php": {
    "install": true,
    "versions": [
      "8.1"
    ],
    "default_version": "8.1",
    "enable_service": true
  },
  "blackfire": {
    "install": false,
    "collector": "https://blackfire.io",
    "log_level": 1,
    "server_id": "",
    "server_token": "",
    "socket": "unix:///var/run/blackfire/agent.sock",
    "enable_service": true
  },
  "nodejs": {
    "install": true,
    "version": "18",
    "install_yarn": true
  },
  "golang": {
    "install": false,
    "version": "1.19"
  },
  "nginx": {
    "install": true,
    "compile_version": "",
    "enable_service": true
  },
  "memcached": {
    "install": true,
    "enable_service": true
  },
  "redis": {
    "install": true,
    "enable_service": true
  },
  "import_hosts_files": [],
  "extra_ansible_tasks": {
    "task_files": [],
    "variable_file": []
  },
  "extra_service_names": []
}
```

### Variable definition

| Variable Name           | Type     | Description                                                                     |
|-------------------------|----------|---------------------------------------------------------------------------------|
| php                     | object   | [PHP configuration](#php-configuration)                                         |
| blackfire               | object   | [Blackfire configuration](#blackfire-configuration)                             |
| nodejs                  | object   | [NodeJS configuration](#nodejs-configuration)                                   |
| golang                  | object   | [GoLang configuration](#golang-configuration)                                   |
| nginx                   | object   | [Nginx configuration](#nginx-configuration)                                     |
| memcached               | object   | [Memcached configuration](#memcached-configuration)                             |
| redis                   | object   | [Redis configuration](#redis-configuration)                                     |
| import_hosts_files      | []object | [Import hosts configuration](#import-hosts-configuration)                       |
| extra_ansible_playbooks | object   | [Extra Ansible Playbooks configuration](#extra-ansible-playbooks-configuration) |
| extra_service_names     | []string | [Extra service names configuration](#extra-service-names-configuration)         |

### PHP configuration

| Variable Name   | Type     | Description                   | Default     |
|-----------------|----------|-------------------------------|-------------|
| install         | bool     | Install PHP or not            | `true`      |
| versions        | []string | Setup PHP versions to install | `["8.1"]`   |
| default_version | string   | Setup a default PHP version   | `"8.1"`     |
| enable_service  | bool     | Enable PHP-FPM at startup     | `true`      |

### Blackfire configuration

| Variable Name  | Type   | Description                                 | Default                                  |
|----------------|--------|---------------------------------------------|------------------------------------------|
| install        | bool   | Install Blackfire or not                    | `false`                                  |
| collector      | string | Setup blackfire-agent config "collector"    | `"https://blackfire.io"`                 |
| log_level      | int    | Setup blackfire-agent config "log-level"    | `1`                                      |
| server_id      | string | Setup blackfire-agent config "server-id"    | `""`                                     |
| server_token   | string | Setup blackfire-agent config "server-token" | `""`                                     |
| socket         | string | Setup blackfire-agent config "socket"       | `"unix:///var/run/blackfire/agent.sock"` |
| enable_service | bool   | Enable blackfire-agent at startup           | `true`                                   |

### NodeJS configuration

| Variable Name  | Type   | Description                     | Default |
|----------------|--------|---------------------------------|---------|
| install        | bool   | Install NodeJS or not           | `true`  |
| version        | string | Setup NodeJS version to install | `"18"`  |
| install_yarn   | bool   | Install `yarn` or not           | `true`  |

### Golang configuration

| Variable Name  | Type   | Description                     | Default  |
|----------------|--------|---------------------------------|----------|
| install        | bool   | Install Golang or not           | `false`  |
| version        | string | Setup Golang version to install | `"1.19"` |

### Nginx configuration

| Variable Name   | Type   | Description                                                                         | Default |
|-----------------|--------|-------------------------------------------------------------------------------------|---------|
| install         | bool   | Install Nginx or not                                                                | `true`  |
| compile_version | string | Set a specific version number to compile Nginx from source, e.g. `1.16.1`, `1.18.0` | `""`    |
| enable_service  | bool   | Enable Nginx at startup                                                             | `true`  |

GearBox will install Nginx from package manager (e.g. `apt`) by default, and install the latest version of Nginx.
If you want to use a specific Nginx version, set up with `compile_version` variable.
And GearBox will install specific Nginx version that compiled from source code.

- Valid `compile_version` values can be found at [https://nginx.org/en/download.html](https://nginx.org/en/download.html).
- The recommended `compile_version` value in WSL is `1.16.1`.

### Memcached configuration

| Variable Name  | Type | Description                 | Default |
|----------------|------|-----------------------------|---------|
| install        | bool | Install Memcached or not    | `true`  |
| enable_service | bool | Enable Memcached at startup | `true`  |

### Redis configuration

| Variable Name  | Type | Description             | Default |
|----------------|------|-------------------------|---------|
| install        | bool | Install Redis or not    | `true`  |
| enable_service | bool | Enable Redis at startup | `true`  |

### Import hosts configuration

`import_hosts_files` configuration is an array of objects, each object definition as below:

| Variable Name | Type   | Description                      | Default |
|---------------|--------|----------------------------------|---------|
| name          | string | The block name of `<block_name>` |         |
| path          | string | The hosts file path to import    |         |

Use this configuration to import custom hosts to system `hosts` files.

- In Linux, will import to `/etc/hosts` file
- In Windows, will import to `C:\Windows\System32\drivers\etc\hosts` file

Example content of these hosts files (e.g. `/path/to/your/hosts.txt`):

```
127.0.0.1 foobar.local.dev
127.0.0.1 foo.local.dev
192.168.0.1 bar.local.dev
```

Example of the imported hosts file `/etc/hosts`:

```
127.0.0.1 localhost
192.168.0.100 example.dev

##>>> INSERTED BY GEARBOX ## block1 ## START >>>##
127.0.0.1 foobar.local.dev
127.0.0.1 foo.local.dev
192.168.0.1 bar.local.dev
##<<< INSERTED BY GEARBOX ## block1 ## END   <<<##
##>>> INSERTED BY GEARBOX ## block2 ## START >>>##
127.0.0.1 foobaz.local.dev
##<<< INSERTED BY GEARBOX ## block2 ## END   <<<##
```

Imported hosts will be placed in a block.
Each imported block is started with a line `##>>> INSERTED BY GEARBOX ## <block_name> ## START >>>##`,
and is ended with a line `##<<< INSERTED BY GEARBOX ## <block_name> ## END   <<<##`.
`<block_name>` is defined in the `name` property, which used to distinguish different imported files.

### Extra Ansible Playbooks configuration

`extra_ansible_playbooks` setup extra ansible playbook files and variable files to run after the main installation tasks.

| Variable Name  | Type     | Description                                                                                      | Default |
|----------------|----------|--------------------------------------------------------------------------------------------------|---------|
| playbook_files | []string | An array of file paths that contains Ansible Playbooks                                           | `[]`    |
| variable_files | []string | An array of file paths that contains variables using in Ansible Playbooks (format: JSON or YAML) | `[]`    |

- In Linux, file path with the form `/path/to/file.yaml`
- In Windows, file path with the form `C:\\path\\to\\file.yaml`, or with WSL file path (e.g. `/mnt/c/path/to/file.yaml`)

Example playbook file (e.g. `/path/to/playbook.yaml`):

```yaml
---
- name: Debug example
  hosts: localhost
  become: true
  vars:
    foo: 'foo'
    bar: '{{ foobar }}'
  tasks:
    - name: Debug example
      debug:
        msg: "Ansible debug example {{ foo }}"
    - name: Debug example with variable "bar"
      debug:
        msg: '{{ bar }}'
```

Example variable file with JSON format (e.g. `/path/to/variables.json`):

```json
{
  "foobar": "This is a variable file example.",
  "list": ["foo", "bar", "foobar"]
}
```

Example variable file with YAML format (e.g. `/path/to/variables.yaml`):

```yaml
---
# Some comments...
foobar: This is a variable file example.
list:
  - "foo"
  - "bar"
  - "foobar"
```

## Extra service names configuration

Add extra service names to `gearbox-service` command.

For example: `{"extra_service_names": ["cron", "ssh"]}`, then you can use commands: `gearbox-service cron restart` or `gearbox-service ssh restart`.
