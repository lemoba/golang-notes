package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	var b int32

	a = 15
	b = a + a
	b = b + 5
	fmt.Println(a, b)
	fmt.Printf("%d", unsafe.Sizeof(a))

}
