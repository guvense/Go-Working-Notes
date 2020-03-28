package main

import (
	"calendar/date"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Event struct {
	title     string
	date.Date // ------------------> > > Embedded using an anonymous field
}

func (e *Event) SetTitle(title string) error {

	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string {
	return e.title
}
func main() {
	event := Event{}
	event.SetYear(10)
	event.SetTitle("deneem")
	fmt.Println(event)
}
