package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = " hello world "
	var s1 string = "@hello@world@"
	fmt.Println(strings.TrimSpace(s))  //去除开头和结尾的空白
	fmt.Println(strings.Trim(s1, "@")) //去除开头和结尾的指定字符
}
