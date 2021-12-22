package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...) // 可变参数
	fmt.Printf("The minimum in the slice is: %d\n", x)
}

// 可变参数
func min(s ...int) int {
	if len(s) == 1 {
		return 0
	}

	min := s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
