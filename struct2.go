package main

import "fmt"

type subscriber struct {
	name    string
	surname string
}

func main() {

	var sub subscriber

	setSub(&sub, "guven", "seckin")

	fmt.Printf("%#v \n", sub)

}

func setSub(sub *subscriber, name string, surname string) subscriber {

	sub.name = name
	sub.surname = surname
	return *sub

}
