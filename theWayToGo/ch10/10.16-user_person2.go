package main

import (
	"fmt"

	"github.com/lemoba/golang-notes/theWayToGo/ch10/person"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("ranen")
	fmt.Println(p.FirstName())
}
