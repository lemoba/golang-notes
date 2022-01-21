package main

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(id string) {
	b.id = id
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float32
}

func main() {
	idjb := Base{"007"}
	jb := Person{idjb, "James", "Bond"}
	e := &Employee{jb, 10000.}
	fmt.Printf("ID of our hero: %v\n", e.Id())
	e.SetId("007b")
	fmt.Printf("The new ID of our hero: %v\n", e.Id())
}
