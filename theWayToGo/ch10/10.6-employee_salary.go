package main

import "fmt"

type employee struct {
	salary float32
}

func (e *employee) giveRaise(pct float32) {
	e.salary *= (1 + pct/100)
}

func main() {
	ep := new(employee)
	ep.salary = 1000
	ep.giveRaise(40)
	fmt.Printf("Employee now makes %.2f\n", ep.salary)
}
