package main

import "fmt"

type Animal struct {
	Name    string
	Species string
	Age     uint
	Sound   string
}

func ModifySound(animal Animal) string {
	animal.Sound = "Roar"
	return animal.Sound
}
func InfoAnimal(animal Animal) int {
	fmt.Println(animal.Name)
	fmt.Println(animal.Species)
	fmt.Println(animal.Age)
	fmt.Println(animal.Sound)
	return 0
}
