package main

import (
	"fmt"
)

func main() {
	val := 256
	// array of integers that is 256 length
	var x [256]int
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < val; i++ {
		x[i] = i
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b \n", v, v, v)
		if i > 50 {
			break
		}
	}
}
