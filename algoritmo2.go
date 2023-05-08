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
func main() {
	person := Person{name: "Marcola", age: 18, address: Address{street: "Primeira Avenida", number: 502, city: "Brasília", state: "DF"}}
	fmt.Print(CompleteAddress(person).address)
}
