package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno",
	}
	for key, val := range myGreeting {
		fmt.Println(key, " _ ", val)
	}
}
