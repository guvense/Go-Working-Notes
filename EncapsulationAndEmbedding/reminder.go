package main

import (
	"errors"
	"fmt"
	"log"
)

// prevent to access these variables they should begin with not capital letter
// Put these codes another package
type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) (*Date, error) {

	if year < 1 {
		return nil, errors.New("invalid year")
	}

	d.year = year
	return d, nil
}

func (d *Date) SetMonth(month int) (*Date, error) {

	if month > 12 || month <= 0 {
		return nil, errors.New("invalid month")
	}
	d.month = month
	return d, nil
}

func (d *Date) SetDay(day int) (*Date, error) {

	if day > 31 || day <= 0 {
		return nil, errors.New("invalid day")
	}
	d.day = day
	return d, nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func main() {
	// struct literal
	date := Date{}

	dat, err := date.SetDay(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dat.Day())

	fmt.Println(date)
}
