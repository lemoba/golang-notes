package main

import "fmt"

func main() {
	for i := 1; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	s := "G"
	for i := 0; i < 25; i++ {
		fmt.Printf(s)
		s += "G"
		fmt.Println()
	}
}
