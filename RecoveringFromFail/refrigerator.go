package main

import "fmt"

func main() {

	var r Refrigerator
	r.addObject("kdsdk")
	r.printAll()

}

type Refrigerator []string

func (r *Refrigerator) addObject(obj string) {

	fmt.Println("Ref opening")
	defer fmt.Println("Ref closing")
	*r = append(*r, obj)

}

func (r Refrigerator) printAll() {

	for _, elem := range r {

		fmt.Println(elem)
	}
}

func (r *Refrigerator) clearAll() {

	*r = nil
}
