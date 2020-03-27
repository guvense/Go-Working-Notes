package main

import "fmt"

func main() {

	var mystruct struct {
		field1 string
		field2 int
	}

	mystruct.field1 = "selam"
	mystruct.field2 = 10
	fmt.Printf("%#v\n", mystruct)

	// var subscriber struct {
	// 	name   string
	// 	rate   float64
	// 	active bool
	// }

	// subscriber.name = "Aman Singh"
	// subscriber.rate = 4.99
	// subscriber.active = true

	type subscriber struct {
		name   string
		rate   float64
		active bool
	}

	sub := subscriber{
		name:   "guven",
		rate:   16.5,
		active: true,
	}
	fmt.Printf("%#v\n", sub)

}
