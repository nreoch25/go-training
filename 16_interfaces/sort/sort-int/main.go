package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	//fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	//fmt.Println(n)

	// reverse
	j := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(j)
	sort.Sort(sort.Reverse(sort.IntSlice(j)))
	fmt.Println(j)
}
