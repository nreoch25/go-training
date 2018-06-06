package main

import (
	"fmt"
)

func zero(z *int) {
	// z is a pointer to the x memory address
	fmt.Println(z)
	// change the value of z
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	// x is now 0
	fmt.Println(x)
}
