package main

import "fmt"

func main() {
	var mapList map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1, "two": 2, "three": 3}
	mapCreated := make(map[string]float32)
	mapAssigned = mapList // 引用

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])      // 1
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"]) // 3.14159
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])     // 3
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])      // 0

}
