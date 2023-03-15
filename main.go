package main

import "fmt"

func main() {
	type Animal struct {
		name string
		age int8	
	}

	animal := Animal{
		name: "Bubbles",
		age: 3,
	}

	fmt.Println(animal)

	type Dog struct {
		numberOfLegs int8
		animal Animal
	}

	dog := Dog{
		numberOfLegs: 4,
		animal: animal,	
	}

	/*
	The statement below doesnt change the value of the animal
	Printing both the dog and animal after the mutation verifies this
	This bevahoiur is so because when we passs a struct to another struct,
	golang makes a copy therefore the struct that gets mutated is a copy
	*/
	dog.animal.name = "Odi"

	fmt.Println(dog)
	fmt.Println(animal)

	/*
	In order to use main reference to the struct, we'll need to use
	pointers
	*/

	type Dog2 struct {
		numberOfLegs int8
		animal *Animal
	}

	dog2 := Dog2{
		numberOfLegs: 4,
		animal: &animal,
	}

	fmt.Println(dog2)//in place of the animal, the address is instead printed
	fmt.Println(*dog2.animal)

	//lets mutate the dog name
	dog2.animal.name = "Bosco"

	fmt.Println(*dog2.animal)
	fmt.Println(animal)
	/*
	The value returned from the last two print statements
	are identical, demonstrating how we can mutate the value
	of a parent struct from a child struct
	*/
}