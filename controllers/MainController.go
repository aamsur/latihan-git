package controllers

import (
	"fmt"
	"reflect"
)

func Calculate(numbers ...interface{}) float64 {
	fmt.Println(len(numbers))

	if (len(numbers) == 0) {
		return 0
	}

	for _, number := range numbers {
		fmt.Println(reflect.TypeOf(number))
	}
	
	return 666
}