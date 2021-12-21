package main

import "fmt"

func main() {
	str := "Go is a beatiful language!"

	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c\n", pos, char)
	}
	fmt.Println()

	str2 := "Chinese: 韩国语"
	fmt.Printf("The length of str2 is: %d\n", len(str2))

	for pos, char := range str2 {
		fmt.Printf("Character %c starts at byte postion %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune char bytes")
	for index, rune := range str2 {
		fmt.Printf("%2d      %d      %U '%c' %X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
