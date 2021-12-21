package main

import "fmt"

func main() {
	var a [5]int

	fmt.Println(a) //初始化为0
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	// for i, j := range twoD {
	// 	for k, _ := range j { // _ 省略
	// 		twoD[i][k] = i + k
	// 	}
	// }
	fmt.Println("2d:", twoD)
}
