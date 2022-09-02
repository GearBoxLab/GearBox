#!/usr/bin/env bash

GEARBOX_LATEST_VERSION="$(curl -s 'https://raw.githubusercontent.com/GearBoxLab/GearBox/master/version/latest')"
GEARBOX_DOWNLOAD_URL=https://github.com/GearBoxLab/GearBox/releases/download/$GEARBOX_LATEST_VERSION/gearbox-$GEARBOX_LATEST_VERSION-linux-amd64
DOWNLOAD_GEARBOX=1

if [[ -f /usr/local/bin/gearbox ]]; then
  INSTALLED_GEARBOX_VERSION="$(/usr/local/bin/gearbox version --no-ansi)"

  if [[ $INSTALLED_GEARBOX_VERSION == *"$GEARBOX_LATEST_VERSION"* ]]; then
    DOWNLOAD_GEARBOX=0
  fi
fi

if [[ 1 == $DOWNLOAD_GEARBOX ]]; then
  sudo curl -sLo /usr/local/bin/gearbox $GEARBOX_DOWNLOAD_URL
  sudo chmod 755 /usr/local/bin/gearbox
fi

/usr/local/bin/gearbox help
