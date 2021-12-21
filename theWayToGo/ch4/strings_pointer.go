package main

import "fmt"

// 你不能获取字面量或常量的地址，例如：
// const i = 5
// ptr := &i //error: cannot take the address of i
// ptr2 := &10 //error: cannot take the address of 10

func main() {
	s := "good bye"
	var p *string = &s
	*p = "giao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}
