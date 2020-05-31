package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {

	return string(c) + "coffee pot"
}

func main() {
	co := CoffeePot("hello")
	fmt.Println(co)
}

// print use String interface method like toString method in java
