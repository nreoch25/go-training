package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	// variadic arguments dots at end
	n := average(data...)
	fmt.Println(n)
}

// variadic params dots at front
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
