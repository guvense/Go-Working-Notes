package main

import "fmt"

type Liters float64
type Gallons float64

// no overloading

func main() {

	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.)
	busFuel = Liters(240.)

	fmt.Println(carFuel, busFuel)

	fmt.Println(carFuel.toLitter().toGallon().toLitter(), busFuel.toGallon())

	busFuel.set(120)

	fmt.Println(carFuel, busFuel)

}

func (l Liters) toGallon() Gallons {

	return Gallons(l * 0.264)
}

func (g Gallons) toLitter() Liters {

	return Liters(g * 3.785)
}

func (l *Liters) set(num float64) {
	*l = Liters(num)
}
