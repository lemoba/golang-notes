package main

import "fmt"

// 但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败
func main() {
	a := 1
	goto TAGGET // complie error
	b := 9
TAGGET:
	b += a
	fmt.Printf("a is %v *** b is %v\n", a, b)
}
