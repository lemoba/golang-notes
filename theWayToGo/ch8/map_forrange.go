package main

import "fmt"

func main() {
	map1 := make(map[int]float32) // key是无序的
	map1[0] = 1.0
	map1[1] = 1.0
	map1[2] = 1.0
	map1[3] = 1.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
}
