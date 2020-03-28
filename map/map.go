package main

import (
	"fmt"
)

func main() {

	// Creating map

	var ranks map[string]int
	ranks = make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true

	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])

	// Map literals
	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap["a"])

	var words map[string]string
	fmt.Printf("%#v\n", words["deneme"])

	// if only declaration we get panic
	//words["selam"] = "hi"

	// To prevent panic use
	words := make(map[string]string)
	words["selam"] = "hi"

}
