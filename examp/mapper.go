package main

import (
	"fmt"
	"magazine"
	"reader/stringreader"
	"strings"
)

func main() {

	array := stringreader.GetStrings("people.txt")

	var people []magazine.Person

	for _, elem := range array {

		s := strings.Split(elem, ":")
		person := magazine.CreatePerson(s[0], s[1])

		people = append(people, *person)

	}

	for i, elem := range people {

		fmt.Println(i, elem)
	}

}
