package main

import (
	"fmt"
)

func calmDown() {

	p := recover()

	err, ok := p.(error)

	if ok {
		fmt.Println(err.Error())
	}

}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	defer calmDown()

	err := ComedyError("Keep slience")

	panic(err)
}
