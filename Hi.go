package main

import (
	"fmt"
)

// Go is a pass-by-value language

func main() {

	myArray := []int{1, 2, 3, 4, 5}

	pointerAddressList := getAddress(myArray)

	printValueFromAdress(pointerAddressList)

	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 10
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

	testNumber := 2
	fmt.Println("Before: ", testNumber)
	fmt.Println("Address first: ", &testNumber)
	changePlz(&testNumber)
	fmt.Println("After: ", testNumber)

}

func getAddress(myArray []int) []*int {

	pointerArray := []*int{}

	for _, elem := range myArray {

		pointerArray = append(pointerArray, &elem)
	}

	return pointerArray
}

func printValueFromAdress(arr []*int) {

	for _, elem := range arr {

		fmt.Println(elem, *elem)
	}
}

func changePlz(val *int) {

	fmt.Println("Address second: ", val)
	*val *= 2
}
