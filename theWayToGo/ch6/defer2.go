package main

import "fmt"

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func main() {
	a()
	f()
}
