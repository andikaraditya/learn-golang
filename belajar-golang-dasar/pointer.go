package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr." + man.Name
}

func main() {
	address1 := Address{"Jogja", "DIY", "Indonesia"}

	// Pass by value
	address2 := address1
	address2.City = "Gunungkidul"
	fmt.Println(address1)
	fmt.Println(address2)

	// Pass by reference
	address3 := &address1
	address3.City = "Sleman"
	fmt.Println(address1)
	fmt.Println(address3)

	// to another pointer
	// address3 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address3)

	// to new value
	*address3 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address3)

	// new
	address4 := new(Address)
	address5 := address4

	address5.City = "Surabaya"
	fmt.Println(address4)
	fmt.Println(address5)

	// change pointer inside function
	address6 := &Address{"New York City", "New York", "USA"}

	changeCountryToIndonesia(address6)

	fmt.Println(address6)

	john := Man{"John"}

	john.married()

	fmt.Println(john)
}
