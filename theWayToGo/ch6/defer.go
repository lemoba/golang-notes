package main

import "fmt"

// defer 推迟执行 栈结构 先进后出
func main() {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")

	defer function2()
	defer function3()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

func function3() {
	fmt.Printf("Function3: Deferred until the end of the calling function!\n")
}
