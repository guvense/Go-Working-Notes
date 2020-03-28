package main

import (
	"coordinates"
	"fmt"
	"log"
)

func main() {

	coor := coordinates.Coordinates{}

	err := coor.SetLang(10)
	printAndExist(err)

	fmt.Printf("%#v", coor)
}

func printAndExist(err error) {

	if err != nil {
		log.Fatal(nil)
	}
}
