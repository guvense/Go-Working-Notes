package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "sol", "la", "si"}

	for i := 0; i < len(notes); i++ {

		fmt.Println(i, notes[i])
	}

	for i, elem := range notes {

		fmt.Println(i, elem)
	}

	myArray := []float64{1., 2., -3.}

	avg, err := average(myArray)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("Your avarage is :", avg)

}

// blank identifier

func average(myArray []float64) (float64, error) {

	size := len(myArray)
	sum := 0.
	for _, elem := range myArray {
		if elem < 0 {

			err := errors.New("Numbers can not be negative")
			return 0., err
		}
		sum += elem
	}

	return sum / float64(size), nil
}
