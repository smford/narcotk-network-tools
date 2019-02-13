package main

import (
	//"flag"
	"fmt"
	"log"
	//"github.com/spf13/pflag"
	"github.com/spf13/viper"
	//"net"
	//"os"
	//"sort"
	//"strings"
)

const (
	configfile = "config-wol"
	configpath = "."
)

var (
	hosts []string
)

func init() {
	fmt.Println("init")
	fmt.Println(configfile)
	fmt.Println(configpath)
	viper.AddConfigPath(configpath)
	viper.SetConfigName(configfile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ERROR loading config: ", err)
	}

}

func main() {
	fmt.Println("main")
	for k, v := range viper.GetStringMapString("Hosts") {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	fmt.Println(viper.AllSettings())
	fmt.Println(viper.GetString("something"))
}
