/usr/bin/bash

## Install "GoHugo" and "Make"
apt-get update && apt-get install -y hugo make

## Building Hugo
make build
