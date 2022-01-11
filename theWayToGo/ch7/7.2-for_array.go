package main

import "fmt"

func main() {
	var arr [15]int

	for i := range arr {
		arr[i] = i
	}

	for i, v := range arr {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}
}
