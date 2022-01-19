package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

func (c *Car) Start() {
	fmt.Println("Car is started!")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped!")
}

func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func main() {
	m := &Mercedes{Car{4, nil}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
