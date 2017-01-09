package main

import (
	"fmt"
	"net/http"

	"github.com/aamsur/latihan-git/calculator"
)

// ServerConfig for sever port and host
type (
	ServerConfig struct {
		Port string
		Host string
	}
)

func main() {
	var conf ServerConfig
	conf.Port = ":1234"
	conf.Host = "0.0.0.0"
	var address_served = conf.Host + conf.Port
	http.HandleFunc("/calc", calculator.CalculatorHandler)
	fmt.Println()
	fmt.Printf("Your Application Running on Addrss : %s", address_served)
	fmt.Println()
	http.ListenAndServe(address_served, nil)
}
