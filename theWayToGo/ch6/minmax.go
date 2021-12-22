package main

import "fmt"

func main() {
	var min, max int
	min, max = minMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}

func minMax(a, b int) (max int, min int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}
