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

		myips, err := net.LookupIP(myhost)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, ip := range myips {
			fmt.Println(ip)
		}
	} else {
		fmt.Println("invalid host:", myhost)
		os.Exit(1)
	}
}

func displayHelp() {
	helpmessage := `
forwardlookup [hostname]
`
	fmt.Println(helpmessage)
}
