package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Nigel", "Nathan", "Rob", "Adrian"}
	sort.Strings(s)
	fmt.Println(s)
}
