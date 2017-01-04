package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"github.com/aamsur/latihan-git/resources"
)

func main() {
	var n int

	a := os.Args[1:]

	var x = 214748364721
	var y = 123
	hasil1, _ := resources.Qasico(123, 321)

	//reflection
	fmt.Println(reflect.TypeOf(x), x + y, hasil1)

	n, _ = strconv.Atoi(a[0]);

	first := 0
	second := 1
	next := 0
	c := 0

	for c = 0; c < n; c++ {
		if c <= 1 {
			next = c;
		} else {
			next = first + second;
			first = second;
			second = next;
		}

		fmt.Println(next)
	}
}