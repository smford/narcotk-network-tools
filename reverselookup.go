package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func init() {
	if len(os.Args) != 2 {
		displayHelp()
		os.Exit(1)
	}
}

func main() {
	var ip string = strings.Join(os.Args[1:], "")
	if ValidIP(ip) {

		myhosts, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, host := range myhosts {
			fmt.Println(strings.TrimRight(host, "."))
		}
	} else {
		fmt.Println("invalid ip:", ip)
		os.Exit(1)
	}
}

func ValidIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}

func displayHelp() {
	helpmessage := `
reverseip [ip address]
`
	fmt.Println(helpmessage)
}
