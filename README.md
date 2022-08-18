GearBox
=======

GearBox helps to easily create the same web development environment in Windows(WSL) and Linux.

## Features

1. Supported OS:
   - Windows 10/11 with WSL enabled. Supported WSL distributions:
       - Ubuntu-18.04
       - Ubuntu-20.04
   - Ubuntu
2. Install packages:
   - PHP (install multiple versions of PHP, v5.6 ~ v8.1)
   - Blackfire
   - NodeJS (include Yarn)
   - GoLang
   - Nginx
   - Memcached
   - Redis

## Windows Requirements

1. Need Windows 10 version 2004 and higher (Build 19041 and higher) or Windows 11.
2. Enable [WSL (Windows Subsystem for Linux)](https://docs.microsoft.com/en-us/windows/wsl/install-manual).
3. Set up WSL default version to 1. (It is recommended to use WSL1 for [better filesystem performance](https://docs.microsoft.com/en-us/windows/wsl/compare-versions#comparing-features))
    
    ```bash
    # Set WSL default version to 1.
    wsl --set-default-version 1
    ```
4. [Install a Linux distribution](https://docs.microsoft.com/en-us/windows/wsl/install).

## Install/Update GearBox in Windows

Run the following command in your terminal to install/update GearBox files.

```bash
bitsadmin.exe /transfer "Download GearBox" https://raw.githubusercontent.com/GearBoxLab/GearBox/master/scripts/install-gearbox-windows.bat %TEMP%\install-gearbox-windows.bat && %TEMP%\install-gearbox-windows.bat && DEL /Q %TEMP%\install-gearbox-windows.bat 
```

## Install/Update GearBox in Linux

```bash
curl -s -o- https://raw.githubusercontent.com/GearBoxLab/GearBox/master/scripts/install-gearbox-linux.sh | bash
```
