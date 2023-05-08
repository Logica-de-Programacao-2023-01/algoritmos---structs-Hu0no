package main

import "fmt"

type Address struct {
	street string
	number uint
	city   string
	state  string
}
type Person struct {
	name    string
	age     uint
	address Address
}

func CompleteAddress(person Person) Person {
	return person
}
