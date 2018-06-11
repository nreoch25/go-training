package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Nigel", "Nathan", "Rob", "Adrian"}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
	// sort.Sort(sort.StringSlice(s))
	// fmt.Println(s)
	// fmt.Printf("%T \n", s)
	// // reverse the slice of strings
	// i := sort.Reverse(sort.StringSlice(s))
	// fmt.Println(i)
	// fmt.Printf("%T \n", i)
	// sort.Sort(i)
	// fmt.Println(s)
}
