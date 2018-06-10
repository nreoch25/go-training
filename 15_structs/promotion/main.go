package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	Person
	first         string
	licenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			first: "Nigel",
			last:  "Reoch",
			age:   39,
		},
		first:         "Double Zero Seven",
		licenseToKill: true,
	}
	fmt.Println(p1.first, p1.Person.first, p1.age, p1.licenseToKill)
}
