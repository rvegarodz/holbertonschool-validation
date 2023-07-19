/usr/bin/bash

## Install Make
apt-get update && apt-get install -y make wget

## Install Hugo
HUGO_VERSION=0.84.0
wget https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.deb
dpkg -i hugo_${HUGO_VERSION}_Linux-64bit.deb
rm hugo_${HUGO_VERSION}_Linux-64bit.deb

## Building Hugo
make build
