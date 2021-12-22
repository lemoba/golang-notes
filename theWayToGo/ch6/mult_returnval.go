package main

import "fmt"

func sumDiffProduct(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}
func sumDiffProductN(a, b int) (s int, d int, p int) {
	s, d, p = a+b, a-b, a*b
	return
}
func main() {
	sum, diff, product := sumDiffProduct(4, 5)
	fmt.Println("Sum:", sum, "| Diff:", diff, "| Product:", product)
	sum, diff, product = sumDiffProductN(4, 5)
	fmt.Println("Sum:", sum, "| Diff:", diff, "| Product:", product)
}
