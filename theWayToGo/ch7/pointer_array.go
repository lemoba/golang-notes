package main

import "fmt"

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func main() {
	var arr [3]int
	f(arr)   // passes a copy of ar
	fp(&arr) // passes a pointer to ar
}
