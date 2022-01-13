package main

import "fmt"

func main() {
	printRec(1)
}

func printRec(i int) {
	if i > 10 {
		return
	}
	printRec(i + 1)
	fmt.Printf("%d ", i)
}
