package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Nigel", "Nathan", "Rob", "Adrian"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	// sort.Sort takes an interface and sort.StringSlice(s) converts to an interface
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
