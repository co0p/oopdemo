//
package main

import (
	"fmt"
	"strconv"
)

// Animal contains Feet and Height and noone knows about it's wisdom
type Animal struct {
	Feet   int
	Height int
	wisdom int
}

// String prints a nice message about the animal including information about feet and height
func (animal Animal) String() string {
	return "I have " + strconv.Itoa(animal.Feet) + " Feet and am " + strconv.Itoa(animal.Height) + "cm tall. I am " + strconv.Itoa(animal.wisdom) + " smart.\n"
}

// IncreaseWisdom increases the wisdom of the animal by steps of 1
func (animal *Animal) IncreaseWisdom() {
	animal.wisdom++
}

func main() {
	dog := Animal{Feet: 4, Height: 50}
	fmt.Print(dog) // will use the String() method on the struct Animal with dog as value receiver

	dog.IncreaseWisdom() // modifies the dog's wisdom using pointer receiver
	fmt.Print(dog)

}
