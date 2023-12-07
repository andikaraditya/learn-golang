package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// empty interface = any
func ups() interface{} {
	// return "ups"
	// return 5
	return true
}

func main() {
	person := Person{Name: "Budi"}
	animal := Animal{Name: "Giraffe"}
	SayHello(person)
	SayHello(animal)
	fmt.Println(ups())
}
