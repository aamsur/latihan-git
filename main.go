package main

import (
	"net/http"
	"fmt"
	"github.com/aamsur/latihan-git/calculator"
)

type(
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
	fmt.Printf("Your Application Running on Address : %s", address_served)
	fmt.Println()

	http.ListenAndServe(address_served, nil)
}