#!/usr/bin/env bash

GEARBOX_LATEST_VERSION="$(curl -s 'https://raw.githubusercontent.com/GearBoxLab/GearBox/master/version/latest')"
GEARBOX_DOWNLOAD_URL=https://github.com/GearBoxLab/GearBox/releases/download/$GEARBOX_LATEST_VERSION/gearbox-$GEARBOX_LATEST_VERSION-linux-amd64

sudo curl -sLo /usr/local/bin/gearbox $GEARBOX_DOWNLOAD_URL
sudo chmod 755 /usr/local/bin/gearbox

/usr/local/bin/gearbox help
