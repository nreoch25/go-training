package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	shadow := dog{animal{"woof"}, true}
	// create a slice of any types using empty interface
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)
}
