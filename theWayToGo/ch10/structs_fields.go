package main

import "fmt"

type st struct {
	i   int
	f   float32
	str string
}

func main() {
	s := new(st)
	s.i = 12
	s.f = 2.14
	s.str = "Ranen"

	fmt.Printf("The int is: %d\n", s.i)
	fmt.Printf("The float is: %f\n", s.f)
	fmt.Printf("The string is: %s\n", s.str)
	fmt.Printf("%v\n", s)
}
