package main

import "fmt"

type AS struct {
	x float32
	int
	string
}

func main() {
	as := AS{3.14, 7, "hello"}
	fmt.Println(as.x, as.int, as.string)
	fmt.Println(as)
}
