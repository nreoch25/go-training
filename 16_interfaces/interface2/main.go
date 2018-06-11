package main

import (
	"fmt"
	"math"
)

// square implements the shape interface
// user defined types are concrete types
type square struct {
	side float64
}

// circle implements the shape interface
// user defined types are concrete types
type circle struct {
	radius float64
}

// for a type to implement this interface
// it needs to have a method of area() float64
// square and circle have that method
// There's no concrete behaviour of type interface
// without the implementation of the stored user defined value
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Interfaces let us write functions that are more
// flexible and adaptable and are not tied to the
// details of one particular implementation
func info(z shape) {
	fmt.Println(z)
	// Method calls against an interface value
	// are polymorphic in nature
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
