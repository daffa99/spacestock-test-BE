package main

import (
	"fmt"
)

type person struct {
	name   string
	gender string
	age    int
}

func NewPerson() person {
	p := person{}
	return p
}

func (p person) Name(name string) person {
	p.name = name
	return p
}

func (p person) Gender(gender string) person {
	p.gender = gender
	return p
}

func (p person) Age(age int) string {
	p.age = age
	descFormat := fmt.Sprintf("%s, %s, %d", p.name, p.gender, p.age)
	return descFormat
}

func main() {
	jon := NewPerson().Name("Jon Snow").Gender("Male").Age(24)
	fmt.Println(jon)
}
