package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := GetFloats(os.Args[1])

	if err != nil {

		fmt.Println(err)
	}

	for _, elem := range num {

		fmt.Println(elem)
	}

}

func OpenFile(filename string) (*os.File, error) {

	fmt.Println("Opening", filename)
	return os.Open(filename)
}

func CloseFile(file *os.File) {

	fmt.Println("Closing File")
	file.Close()
}

func GetFloats(filename string) ([]float64, error) {

	var numbers []float64

	file, err := OpenFile(filename)

	if err != nil {
		return nil, err
	}

	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil

}
