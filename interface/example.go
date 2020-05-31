package main

import (
	"fmt"
	"strconv"
)

type CreditCard struct {
	number string
	user
	budget float64
}

type BankAccount struct {
	iban   string
	budget float64
}

type user struct {
	name string
}

// errors

type NoError string

func (c NoError) Error() string {
	return ""
}

type CreditCardError string

func (c CreditCardError) Error() string {
	return string(c)
}

type PaymentError string

func (c PaymentError) Error() string {
	return string(c)
}

type BankAccountError string

func (c BankAccountError) Error() string {
	return string(c)
}

// Getter Setter functions

func (c CreditCard) Number() string {
	return c.number
}

func (c *CreditCard) SetNumber(num string) error {

	if _, err := strconv.Atoi(num); err != nil {
		return CreditCardError("value is not number")
	}

	if size := len([]rune(num)); size != 16 {
		return CreditCardError("value should be 16")
	}

	c.number = num

	return nil

}

func (c CreditCard) Name() string {
	return c.name
}

func (c *CreditCard) SetName(name string) error {

	if size := len([]rune(name)); size >= 16 || size < 4 {
		return CreditCardError("name shoul be between 16 and 3 char")
	}

	c.name = name
	return nil

}

func (c CreditCard) Budget() float64 {
	return c.budget
}

func (c *CreditCard) SetBudget(budget float64) error {

	if budget < 0 {
		return CreditCardError("Budget should be greater that zero")
	}

	c.budget = budget
	return nil

}

func (b BankAccount) Iban() string {

	return b.iban
}

func (b *BankAccount) SetIban(iban string) error {

	if _, err := strconv.Atoi(iban); err != nil {
		return BankAccountError("iban is not number")
	}

	if size := len([]rune(iban)); size != 16 {
		return BankAccountError("iban should be 16")
	}

	b.iban = iban
	return nil

}

func (b BankAccount) Budget() float64 {
	return b.budget
}

func (b *BankAccount) SetBudget(budget float64) error {

	if budget < 0 {
		return BankAccountError("Budget should be greater that zero")
	}

	b.budget = budget
	return nil

}

type Payment interface {
	Pay(price float64) (float64, error)
}

func (c *CreditCard) Pay(price float64) (float64, error) {

	const discount float64 = 10.0

	n := price - ((price * discount) / 100)

	if n > c.budget {
		return 0, PaymentError("Balance not enough")
	}

	c.budget = c.budget - n

	return c.budget, nil

}

func (b *BankAccount) Pay(price float64) (float64, error) {

	const discount float64 = 5.0

	n := price - ((price * discount) / 100)
	fmt.Println(n)
	if n > b.budget {
		return 0, PaymentError("Balance not enough")
	}

	b.budget = b.budget - n

	return b.budget, nil

}

func main() {

	var pay Payment

	pay = &CreditCard{}

	p, ok := pay.(*CreditCard)

	if ok {
		p.SetBudget(100.)
		p.SetName("Guven")
		p.SetNumber("1234123412341234")
	}

	fmt.Println("%#v", p)
	bugdet, err := p.Pay(10)

	fmt.Println(err)
	fmt.Println(bugdet)

}
