package main

import (
	"fmt"
)

func main() {

	slice := []string{"a", "b"}
	a := &slice[0]
	fmt.Println(&a)
	fmt.Println(slice, len(slice))
	slice = append(slice, "c", "a", "a")
	a = &slice[0]
	fmt.Println(&a)
	fmt.Println(slice, len(slice))

	floatSlice := make([]float64, 2)
	fmt.Println(floatSlice)
	floatSlice = append(floatSlice, 1)
	fmt.Println(floatSlice)

}
