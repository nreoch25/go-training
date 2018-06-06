package main

import (
	"fmt"
)

func main() {
	var small int
	var large int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&large)
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)
	fmt.Println(large / small)
}
