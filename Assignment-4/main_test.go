package main

import (
	"testing"
)

var (
	p Person
)

func TestValidateValid(t *testing.T) {
	p.Name = "Jon Snow"
	p.Gender = "Male"
	p.Age = 24

	t.Logf("Validate Name: %s, Gender: %s, Age: %d", p.Name, p.Gender, p.Age)

	if p.Validate() != nil {
		t.Errorf("Incorrect! It should've return nil")
	}
}

func TestValidateInvalidName(t *testing.T) {
	p.Name = ""
	p.Gender = "Male"
	p.Age = 24
	t.Logf("Validate Name: %s, Gender: %s, Age: %d", p.Name, p.Gender, p.Age)

	if p.Validate() == nil {
		t.Errorf("Incorrect! Name cannot be empty")
	}
}

func TestValidateInvalidGender(t *testing.T) {
	p.Name = "Jon Snow"
	p.Gender = "Custom"
	p.Age = 24
	t.Logf("Validate Name: %s, Gender: %s, Age: %d", p.Name, p.Gender, p.Age)

	if p.Validate() == nil {
		t.Errorf("Incorrect! Gender is either Male or Female")
	}
}

func TestValidateInvalidAge(t *testing.T) {
	p.Name = "Jon Snow"
	p.Gender = "Male"
	p.Age = -1
	t.Logf("Validate Name: %s, Gender: %s, Age: %d", p.Name, p.Gender, p.Age)

	if p.Validate() == nil {
		t.Errorf("Incorrect! There is no such thing as negative age")
	}
}
