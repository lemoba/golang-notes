package main

import (
	"fmt"
	"strings"
)

// Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new：
// strings.Replace(str, old, new, n) string

func main() {
	var s = "hello world"
	fmt.Println(strings.Replace(s, "hel", "wol", 3))
	fmt.Println(strings.Replace(s, "wollo", "world", -1))
}
