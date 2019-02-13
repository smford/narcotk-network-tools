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
	var myhost string = strings.Join(os.Args[1:], "")
	if true {

		mydnsservers, err := net.LookupNS(myhost)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, nsserver := range mydnsservers {
			fmt.Println(strings.TrimRight(nsserver.Host, "."))
		}
	} else {
		fmt.Println("invalid host:", myhost)
		os.Exit(1)
	}
}

func displayHelp() {
	helpmessage := `
dnslookup [hostname]
`
	fmt.Println(helpmessage)
}
