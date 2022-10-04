@ECHO OFF

SETLOCAL ENABLEDELAYEDEXPANSION

SET VERSION_FILE_PATH=%TEMP%\GearBox-latest-version.txt
SET REFRESHENV_PATH=%TEMP%\RefreshEnv.cmd
SET DOWNLOAD_GEARBOX=1

CALL bitsadmin.exe /transfer "Get GearBox latest version" https://raw.githubusercontent.com/GearBoxLab/GearBox/master/version/latest %VERSION_FILE_PATH%
SET /p GEARBOX_LATEST_VERSION=<%VERSION_FILE_PATH%
ECHO GearBox latest version is %GEARBOX_LATEST_VERSION%

SET GEARBOX_DOWNLOAD_URL=https://github.com/GearBoxLab/GearBox/releases/download/%GEARBOX_LATEST_VERSION%/gearbox-%GEARBOX_LATEST_VERSION%-windows-amd64.exe

IF NOT EXIST "%USERPROFILE%\.gearbox\bin" MKDIR "%USERPROFILE%\.gearbox\bin"

IF EXIST "%USERPROFILE%\.gearbox\bin\gearbox.exe" (
  SET INSTALLED_GEARBOX_VERSION=unknown
  FOR /F "tokens=* USEBACKQ" %%A IN (`"%USERPROFILE%\.gearbox\bin\gearbox.exe" version --no-ansi`) DO (
    SET INSTALLED_GEARBOX_VERSION=%%A
  )

  ECHO !INSTALLED_GEARBOX_VERSION!|FINDSTR /c:"%GEARBOX_LATEST_VERSION%" >nul && SET DOWNLOAD_GEARBOX=0
)

IF %DOWNLOAD_GEARBOX% == 1 (
  CALL bitsadmin.exe /transfer "Download gearbox.exe (%GEARBOX_LATEST_VERSION%)" %GEARBOX_DOWNLOAD_URL% %USERPROFILE%\.gearbox\bin\gearbox.exe
  CALL bitsadmin.exe /transfer "Download chocolatey's RefreshEnv.cmd" https://raw.githubusercontent.com/chocolatey/choco/f924d47fb4177a9a34ff0c2bf995938b5c12800b/src/chocolatey.resources/redirects/RefreshEnv.cmd %REFRESHENV_PATH%

  "%USERPROFILE%\.gearbox\bin\gearbox.exe" init

  ECHO.

  CALL %REFRESHENV_PATH%
  DEL /Q %REFRESHENV_PATH%
)

ECHO.

"%USERPROFILE%\.gearbox\bin\gearbox.exe" help
