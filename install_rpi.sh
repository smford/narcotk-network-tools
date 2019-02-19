#!/usr/bin/env bash

apps=( dnslookup forwardlookup publicip reverselookup )

set -x

mkdir ~/rpi-binaries
cd ~/rpi-binaries

for i in "${apps[@]}"
do
  curl -O -L https://github.com/smford/narcotk-network-tools/raw/master/binaries/rpi/%i
  chmod +x %i
done
