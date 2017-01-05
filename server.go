package main

import (
	"github.com/aamsur/latihan-git/controllers"
	"fmt"
)

func main() {
	//var sl = []int{2,3,4,5,6,7}

	var avg = controllers.Calculate(5,6,7,"delapan")
	//var avg = controllers.Calculate(sl...)
	var msg = fmt.Sprintf("Rata-rata	:	%.2f", avg)
	fmt.Println(msg)
}