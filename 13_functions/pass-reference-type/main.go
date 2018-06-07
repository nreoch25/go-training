package main

import (
	"fmt"
)

func main() {
	m := make([]string, 1, 25)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "hello"
	fmt.Println(z)
}
