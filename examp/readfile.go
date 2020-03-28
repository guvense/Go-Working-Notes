package main

import (
	"fmt"
	"os"
	"reader"
)

func main() {

	// /home/guven/Desktop/GO/data.txt
	myArray := reader.GetFloats(os.Args[1])

	for i, elem := range myArray {

		fmt.Println(i, elem)
	}

}
