package main

import "fmt"

var fib [100]int64

func main() {
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	for i := 0; i < 50; i++ {
		fmt.Println(fib[i])
	}
}
