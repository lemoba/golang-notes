package main

import (
	"fmt"

	"github.com/lemoba/golang-notes/theWayToGo/ch4/trans"
)

var twoPi = 2 * trans.Pi

func main() {
	fmt.Printf("2 * Pi = %g\n", twoPi)
}
