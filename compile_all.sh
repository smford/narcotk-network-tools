#!/usr/bin/env bash

apps=( dnslookup forwardlookup publicip reverselookup )

set -x

for i in "${apps[@]}"
do
  GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o binaries/x86_64/$i $i.go; upx binaries/x86_64/$i
  GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o binaries/rpi/$i $i.go; upx binaries/rpi/$i
  GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o binaries/osx/$i $i.go; upx binaries/osx/$i
done
