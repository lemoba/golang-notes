package main

import "fmt"

func main() {
	//var s int = 3

	if s := 3; s > 3 {
		fmt.Println("up")
	} else if s == 3 {
		fmt.Println("equal 3")
	} else {
		fmt.Println("low")
	}
}
