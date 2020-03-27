package main

import (
	"printer"
	"reader/stringreader"
)

func main() {

	myArray := stringreader.GetStrings("string.txt")

	var names []string
	var counts []int

	for _, elem := range myArray {

		index := getIndex(names, elem)

		if index == -1 {
			names = append(names, elem)
			counts = append(counts, 1)
		} else {

			tmp := counts[index]
			counts[index] = tmp + 1

		}

	}

	printer.ArrayPrinterString(names)
	printer.ArrayPrinterInt(counts)
}

func getIndex(names []string, val string) int {

	for i, elem := range names {

		if elem == val {
			return i
		}
	}

	return -1
}
