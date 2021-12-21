package main

import "fmt"

func main() {
	var i = 5

	//一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关
	var intP *int = &i
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i, &i)
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
