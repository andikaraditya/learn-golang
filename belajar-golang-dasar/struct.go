package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var eko Customer

	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 20
	fmt.Println(eko)

	budi := Customer{
		Name:    "Budi",
		Address: "Indonesia",
		Age:     25,
	}

	fmt.Println(budi)

	budi.sayHello("John")
}
