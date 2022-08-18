@ECHO OFF

IF EXIST %TEMP%\ (
  SET VERSION_FILE_PATH=%TEMP%\GearBox-latest-version.txt
) ELSE (
  SET VERSION_FILE_PATH=%~dp0GearBox-latest-version.txt
)
bitsadmin.exe /transfer "Get GearBox latest version" https://raw.githubusercontent.com/GearBoxLab/GearBox/master/version/latest %VERSION_FILE_PATH%
SET /p GEARBOX_LATEST_VERSION=<%VERSION_FILE_PATH%

SET GEARBOX_DOWNLOAD_URL=https://github.com/GearBoxLab/GearBox/releases/download/%GEARBOX_LATEST_VERSION%/gearbox-%GEARBOX_LATEST_VERSION%-windows-amd64.exe

MKDIR %USERPROFILE%\.gearbox\bin

bitsadmin.exe /transfer "Download gearbox.exe" %GEARBOX_DOWNLOAD_URL% %USERPROFILE%\.gearbox\bin\gearbox.exe

%USERPROFILE%\.gearbox\bin\gearbox.exe init

REFRESHENV

%USERPROFILE%\.gearbox\bin\gearbox.exe help
