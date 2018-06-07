package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T \n", mySlice)

	for i, value := range mySlice {
		fmt.Println(i, " - ", value)
	}

}
