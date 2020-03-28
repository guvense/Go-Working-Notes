package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	argumentstr := os.Args[1:]

	average := getAverage(argumentstr)

	fmt.Println("Average: ", average)

}

func getAverage(arguments []string) float64 {

	var sum float64 = 0

	for _, elem := range arguments {
		argument, err := strconv.ParseFloat(elem, 64)
		printAndExit(err)
		sum += argument

	}

	average := sum / float64(len(arguments))
	return average
}

func printAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
