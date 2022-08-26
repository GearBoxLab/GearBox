CHANGELOG
=========

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
