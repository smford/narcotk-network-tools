package main

// https://www.ipify.org

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(ip[:len(ip)-1]))
}
