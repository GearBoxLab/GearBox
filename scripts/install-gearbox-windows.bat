@ECHO OFF

SET VERSION_FILE_PATH=%TEMP%\GearBox-latest-version.txt
SET REFRESHENV_PATH=%TEMP%\RefreshEnv.cmd

bitsadmin.exe /transfer "Get GearBox latest version" https://raw.githubusercontent.com/GearBoxLab/GearBox/master/version/latest %VERSION_FILE_PATH%
SET /p GEARBOX_LATEST_VERSION=<%VERSION_FILE_PATH%

SET GEARBOX_DOWNLOAD_URL=https://github.com/GearBoxLab/GearBox/releases/download/%GEARBOX_LATEST_VERSION%/gearbox-%GEARBOX_LATEST_VERSION%-windows-amd64.exe

MKDIR %USERPROFILE%\.gearbox\bin

bitsadmin.exe /transfer "Download gearbox.exe" %GEARBOX_DOWNLOAD_URL% %USERPROFILE%\.gearbox\bin\gearbox.exe
bitsadmin.exe /transfer "Download chocolatey's RefreshEnv.cmd" https://raw.githubusercontent.com/chocolatey/choco/f924d47fb4177a9a34ff0c2bf995938b5c12800b/src/chocolatey.resources/redirects/RefreshEnv.cmd %REFRESHENV_PATH%

%USERPROFILE%\.gearbox\bin\gearbox.exe init

ECHO.

CALL %REFRESHENV_PATH%
DEL /Q %REFRESHENV_PATH%

ECHO.

%USERPROFILE%\.gearbox\bin\gearbox.exe help
