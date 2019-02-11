package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"strings"
)

func init() {

}

func main() {
	var ip string = "8.8.8.8"
	myhosts, err := net.LookupAddr(ip)
	if err != nil {
		log.Fatal("Cannot lookup: ", ip, "Error: ", err)
		os.Exit(1)
	}
	fmt.Println("IP:", ip)
	for _, host := range myhosts {
		fmt.Println(strings.TrimRight(host, "."))
	}
}
