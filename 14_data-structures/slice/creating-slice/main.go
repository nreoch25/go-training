package main

import (
	"fmt"
)

func main() {
	// 3 ways to create a slice
	// shorthand
	// must use append
	shorthand := []string{}
	fmt.Printf("%T \n", shorthand)
	// var way
	// zero value of slice - nil
	// must use append
	var varslice []string
	fmt.Printf("%T \n", varslice)
	varslice = append(varslice, "hey")
	fmt.Println(varslice)
	// use make which is considered the best way
	// creates length and array capacity of 35
	makeslice := make([]string, 35)
	makeslice[0] = "hey there"
	makeslice[17] = "hey"
	makeslice[34] = "hey you"
	fmt.Printf("%T \n", makeslice)
	fmt.Println(makeslice)
}

// slice references an underlying array data structure
// use append to add to a slice
