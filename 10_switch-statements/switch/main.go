package main

import (
	"fmt"
)

func main() {
	switch "Mike" {
	case "Daniel":
		fmt.Println("hey dan")
	case "Nigel":
		fmt.Println("hey nigel")
	case "Jenny":
		fmt.Println("hey jenny")
	default:
		fmt.Println("Hey nobody")
	}
}
