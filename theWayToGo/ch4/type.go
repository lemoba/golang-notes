package main

import "fmt"

type TZ int
type Rope string

func main() {
	var a, b TZ = 3, 4
	var s Rope = "hello"
	c := a + b
	fmt.Printf("c has the value: %d\n", c)
	println(s)
}
