package main

import "fmt"

func main() {

	// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
	const (
		a = iota // 0
		b        //1
		c        //2
		d = 5    //3
		e        //4
		f        //5
	)

	fmt.Println(a, b, c, d, e, f)

	// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
	const (
		Apple, Banana     = iota + 1, iota + 2 // 1, 2
		Cherimoya, Durian                      // 2, 3
		Elderberry, Fig                        // 3, 4
	)
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)

	// 使用 iota 结合 位运算 表示资源状态的使用案例
	const (
		Open    = 1 << iota // 0001
		Close               // 0010
		Pending             // 0100
	)
	fmt.Println(Open, Close, Pending)

	const (
		_  = iota             // 使用 _ 忽略不需要的 iota
		KB = 1 << (10 * iota) // 1 << (10*1)
		MB                    // 1 << (10*2)
		GB                    // 1 << (10*3)
		TB                    // 1 << (10*4)
		PB                    // 1 << (10*5)
		EB                    // 1 << (10*6)
		ZB                    // 1 << (10*7)
		YB                    // 1 << (10*8)
	)

	const (
		a1 = 1    // 1
		b1 = iota // 1
		c1 = 4    // 4
		d1 = iota // 3
	)

	fmt.Println(a1, b1, c1, d1)
}
