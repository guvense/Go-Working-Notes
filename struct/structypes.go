package main

import "fmt"

type Address struct {
	location string
}

type Person struct {
	name    string
	Address //anonymous fields
}

func main() {

	p := Person{name: "guven"}
	p.location = "djhdsj" // embedded type can use
	p.Address.location = "kskd"

	fmt.Printf("%#v\n", p)
}
