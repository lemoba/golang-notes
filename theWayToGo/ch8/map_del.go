package main

import "fmt"

func main() {
	m := map[string]int{"name": 1}

	if _, ok := m["name"]; ok {
		fmt.Println(ok)
	}
	fmt.Println(m["name"])
	delete(m, "name")
	fmt.Println(m["name"])
}
