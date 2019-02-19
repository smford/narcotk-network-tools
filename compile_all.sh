#!/usr/bin/env bash

apps=( dnslookup forwardlookup publicip reverselookup )

set -x

for i in "${apps[@]}"
do
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o assets/$i-linux-x86_64 $i.go; upx assets/$i-linux-x86_64
  GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o assets/$i-rpi-arm5 $i.go; upx assets/$i-rpi-arm5
  GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o assets/$i-osx-amd64 $i.go; upx assets/$i-osx-amd64
done
