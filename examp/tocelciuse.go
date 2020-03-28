package main

import (
	"fmt"
	"log"
	"reader"
)

func main() {

	fmt.Println(reader.Deneme)
	fmt.Println("Enter a temparature")
	fahrenheit, err := reader.GetFloat()
	if err != nil {
		log.Fatal(err)

	}
	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
