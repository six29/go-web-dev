package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "Grunt"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func printInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and it has", a.NumberOfLegs(), "legs")
}

func main() {
	dog := Dog{
		Name:  "Summer",
		Breed: "Belgian Malinois",
	}

	printInfo(dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Blue",
		NumberOfTeeth: 33,
	}

	printInfo(gorilla)
}
