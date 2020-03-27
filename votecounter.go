package main

import (
	"fmt"
	"reader/stringreader"
)

func main() {

	myArray := stringreader.GetStrings("string.txt")

	myMap := make(map[string]int)

	for _, elem := range myArray {

		myMap[elem]++
	}

	for key, value := range myMap {

		fmt.Println(key, value)
	}

	// map is an unordered collection of keys and values

}
