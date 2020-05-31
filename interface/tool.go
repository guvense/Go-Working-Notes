package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}

func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {

	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffeePot("LuxBrew")
	device.TurnOn()

	// device.Brew() // Compile time error

	// type assertion
	/*
		When you have a value of a concrete type assigned to a variable with an
		interface type, a type assertion lets you get the concrete type back
	*/

	device = Fan("Windco Breeze")

	c, ok := device.(CoffeePot) // type assertion

	if ok {
		c.Brew()
	} else {
		fmt.Println("Cast error")
	}

}
