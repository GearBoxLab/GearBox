CHANGELOG
=========

## v0.11.2

- Update composer install script

## v0.11.1

- Update installed package versions
  - GoLang: 1.19.7
  - nvm: 0.39.3
  - PHP: 8.2

## v0.11.0

- Fix Nginx compile script
- Update sqlsrv max version for PHP7.4
- Add a new support WSL distribution "Ubuntu-22.04"

## v0.10.1

- Show supported versions in "phpvm" command

## v0.10.0

- Normalize configuration for PHP versions
- Update packages to the newest versions

## v0.9.7

- Update scripts that install Ansible

## v0.9.6

- Fix displaying NodeJS version name and install arguments

## v0.9.5

- Fix displaying NodeJS version name

## v0.9.4

- Change default NodeJS version to "lts"

## v0.9.3

- Fix bug in a Non-English language WSL environment

## v0.9.2

- Fix Ansible tasks

## v0.9.1

- Fix printing duplicated success messages

## v0.9.0

- Add tasks to compile Nginx from source
- Improve Windows UAC process

## v0.8.2

- Fix `import-hosts` command success messages

## v0.8.1

- Fix typo

## v0.8.0

- Fix Windows installer bugs when comparing GearBox versions
- Replace NodeJS installation with `nvm`

## v0.7.0

- Add `extra_service_names` configuration

## v0.6.1

- Do not list install packages with "--only-run-extra-ansible-playbooks" option

## v0.6.0

- Update installer in Windows
- Add a new option `--only-run-extra-ansible-playbooks` to command `gearbox install`

## v0.5.1

- Only download gearbox binary with newer version
- Fix `gearbox-service` script by adding missing "sudo" command

## v0.5.0

- Fix bug that "install-gearbox-windows.bat" stopped after calling `RefreshEnv.cmd`
- Remove debug codes
- Add new options `--sudo-password` and `--yes` to command `gearbox install`

## v0.4.0

- Change config "extra_ansible_tasks" to "extra_ansible_playbooks"

  It is more convenience to use playbook files rather than task files
- Add playbook name mapping for `Ubuntu` distribution to WSL environment

## v0.3.0

- Extend to add multiple variable files (`extra_ansible_tasks.variable_files`) to run extra Ansible Tasks

## v0.2.1

- Fix bug after running "REFRESHENV" command
- Force to use Windows version in WSL distributions
- Convert windows' file path to WSL's file path
- Update README.md for "extra_ansible_tasks" configuration

## v0.2.0

- Add commands:
  - import-hosts: Import hosts with "import_hosts_files" setting
- Add configuration to run extra Ansible tasks
- Update the usage of "phpvm" command
- Add a task to enable blackfire-agent service
- Add configuration document

## v0.1.0

- Add commands:
  - init: Generate default configuration file
  - install: Install packages with Ansible script
- Add install scripts for Windows ([`install-gearbox-windows.bat`](scripts/install-gearbox-windows.bat)) and Linux ([`install-gearbox-linux.sh`](scripts/install-gearbox-linux.sh))
- Add GitHub Action scripts to build binaries, and upload to release
