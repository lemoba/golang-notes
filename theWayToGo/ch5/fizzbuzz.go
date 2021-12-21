package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {
	for i := 1; i < 101; i++ {
		switch {
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}
