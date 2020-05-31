package main

import "fmt"

/*
In Go, an interface is defined as a set of methods that certain values are
expected to have. You can think of an interface as a set of actions you need a
type to be able to perform
*/

// only describe what a value can do: what methods it has.
type myInterface interface {
	Method()
	MethodWithP(float64)
	MethodWithR() string
}

type MyType int

func (m MyType) Method() {

	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithP(f float64) {
	fmt.Println("Method", f)
}

func (m MyType) MethodWithR() string {
	return "Hi"
}

func (my MyType) MethodNotInInterface() {
	fmt.Println("not interface")
}

func main() {

	var value myInterface
	value = MyType(10.)
	value.Method()

}
