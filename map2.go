package main

import (
	"fmt"
)

func main() {

	counters := map[string]int{"a": 3, "b": 0}

	var value int
	var ok bool

	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)

	mapper := map[int]string{1: "bir", 2: "iki", 3: "üç", 4: "dört"}

	deleteFromTo(mapper, 1, 3)

	for _, elem := range mapper {

		fmt.Println(elem)
	}

}

func deleteFromTo(myMap map[int]string, from, to int) {

	for i := from; i <= to; i++ {
		delete(myMap, i)
	}

}
