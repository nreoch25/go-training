package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"
	fmt.Println(myGreeting)
	// shorthand
	myGreeting2 := make(map[string]string)
	myGreeting2["hello"] = "world"
	fmt.Println(myGreeting2)
	// literal
	myGreeting3 := map[string]string{
		"Nigel":  "Reoch",
		"Kristy": "Reoch",
		"Nathan": "Reoch",
		"Hey":    "There",
	}
	// length of a map
	fmt.Println(len(myGreeting3))
	delete(myGreeting3, "Hey")
	fmt.Println(myGreeting3)
	// exists
	if val, exists := myGreeting3["Nathan"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

}
