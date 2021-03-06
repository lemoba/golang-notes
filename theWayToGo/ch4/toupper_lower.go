package main

import (
	"fmt"
	"strings"
)

// ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符：
// strings.ToLower(s) string

// ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符：
// strings.ToUpper(s) string

func main() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
}
