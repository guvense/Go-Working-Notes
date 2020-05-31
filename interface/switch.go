package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toogleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	var t Toogleable = &s // This should be pointer
	s.toggle()
	s.toggle()
}
