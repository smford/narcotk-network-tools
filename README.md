# narcotk-network-utils

| Tool | Use | Example |
| :--- | :--- | :--- |
| forwardlookup | find IPs (ipv4 & ipv6) for a given hostname | `./forwardlookup google.com` |
| reverselookup | find hostname for a given IP address | `./reverselookup 8.8.8.8` |


## Compiling

Linux:
`GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o assets/reverselookup-linux-x86_64 reverselookup.go`


Raspberry Pi
`GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-s -w" -o assets/reverselookup-rpi-arm5 reverselookup.go`

OSX
`GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o assets/reverselookup-osx-amd64 reverselookup.go`
