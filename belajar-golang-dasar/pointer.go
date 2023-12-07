package main

import "fmt"

type Address struct {
	City, Province, Country string
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
}
