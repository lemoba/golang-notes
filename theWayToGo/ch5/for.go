package main

import "fmt"

// go语言只有for
// for 初始化语句; 条件语句; 修饰语句 {}
func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
