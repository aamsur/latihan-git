package main

import (
	"fmt"
)

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = []string{}

	bFruits = append(bFruits, fruits...)

	bFruits[0] = "buah"

	fmt.Println(fruits)
	fmt.Println(bFruits)
}