package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	ipproviderlist = map[string]string{
		"aws":      "https://checkip.amazonaws.com",
		"ipify":    "https://api.ipify.org",
		"my-ip.io": "https://api.my-ip.io/ip",
	}
)

func init() {
	flag.String("provider", "aws", "Provider of your external IP, \"aws\", \"ipify\" or \"my-ip.io\" or \"all\", default = aws")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if !validateipprovider(viper.GetString("provider")) {
		fmt.Printf("--provider %s is not a valid provider\n", viper.GetString("provider"))
		os.Exit(1)
	}

}

func main() {
	if strings.ToLower(viper.GetString("provider")) == "all" {
		for k, v := range ipproviderlist {
			fmt.Printf("%s [%s] %s\n", k, v, getIP(k))
		}
	} else {
		fmt.Println(getIP(viper.GetString("provider")))
	}
}

func getIP(ipprovider string) string {
	ipprovider = strings.ToLower(ipprovider)

	res, err := http.Get(ipproviderlist[ipprovider])
	ip, _ := ioutil.ReadAll(res.Body)

	returnip := string(ip[:len(ip)])
	returnip = strings.TrimSuffix(returnip, "\n")

	if err != nil {
		fmt.Printf("Cannot discern public IP using: %s", ipprovider)
		fmt.Println(err)
		os.Exit(2)
	}

	return returnip
}

func validateipprovider(ipname string) bool {
	ipname = strings.ToLower(ipname)

	if ipname == "all" {
		return true
	}

	for k := range ipproviderlist {
		if k == ipname {
			return true
		}
	}

	return false
}
