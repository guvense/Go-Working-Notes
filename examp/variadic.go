package main

import (
	"fmt"
	"math"
)

func main() {
	value := getMax(1., 3., 3.4)

	fmt.Println(value)

	myArray := inRange(1., 4., 1., 2., 3., 5.)

	for _, elem := range myArray {
		fmt.Println(elem)
	}
}

func getMax(numbers ...float64) float64 {

	max := math.Inf(-1)

	for _, number := range numbers {

		if number > max {
			max = number
		}
	}
	return max

}

func inRange(minVal float64, maxVal float64, numbers ...float64) []float64 {

	var myResult []float64

	for _, elem := range numbers {
		if elem >= minVal && elem <= maxVal {
			myResult = append(myResult, elem)

		}

	}

	return myResult
}
