package main

import (
	"fmt"
	"os"
	"runtime"
)

var (
	v1 string
	v2 int //int8 int16 int32 int64  uint8 uint16 uint32 uint64
	v3 bool
	v4 float64   // float32 float64
	v5 complex64 // complex64 complex128
	v6 byte      // uint8
	v7 rune      // int32
)

func main() {
	var a = "inital"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short" // 只能在函数体内声明
	fmt.Println(f)

	fmt.Println(v1, v2, v3, v4, v5, v6, v7) // 初始化为零值

	var goos string = runtime.GOOS
	fmt.Printf("The opreating system is: %s\n", goos)
	path := os.Getenv("GOPATH")
	fmt.Printf("Path is %s\n", path)
}
