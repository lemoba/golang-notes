package main

import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	fmt.Println(y)

	// NOT OK
	// z := make(Bar) // 编译错误cannot make type Bar
	// (*z).thingOne = "hello"
	// (*z).thingTwo = 1
	// fmt.Println(y)

	// OK
	// x := make(Foo)
	// x["x"] = "goodby"
	// x["y"] = "world"
	// fmt.Println(x)

	// NOT OK
	u := new(Foo) // 运行时错误 panic: assignment to entry in nil map
	(*u)["x"] = "goodby"
	(*u)["y"] = "world"
	fmt.Println(u)

}

// 试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，
// 但是 new() 一个 map 并试图向其填充数据，将会引发运行时错误！
// 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
