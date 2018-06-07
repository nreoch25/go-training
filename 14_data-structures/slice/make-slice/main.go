package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("-------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-------------")

	for i := 0; i < 80; i++ {
		// when length of slice exceeds array capacity
		// the array capacity will get doubled
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
