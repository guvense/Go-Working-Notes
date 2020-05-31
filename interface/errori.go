package main

import (
	"fmt"
)

// error is interface

// type error interface {
// 	Error() string
// }

// We can defice our error type with using Error() function

// predeclared identifier error so it is lower case universe block
type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {

	var err error
	err = ComedyError("Keep slience")
	fmt.Println(err)

}



