package main

import "magazine"

func main() {
	p := magazine.CreatePerson("guven", "seckin")

	p.Print()
	c := magazine.CreateCar(p, "porsche")
	c.Print()

	person := magazine.CreatePerson("nisanco", "coskun")
	person.Deactivate()
	person.Print()

}
