package main

import (
	"fmt"
	"strings"
)

// Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
// strings.Index(s, str string) int

// LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
// strings.LastIndex(s, str string) int
// 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：
// strings.IndexRune(s string, r rune) int

func main() {
	var str string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}
