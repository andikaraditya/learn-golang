package main

import "fmt"

func main() {
	name := "John"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Who are you")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else if length == 5 {
		fmt.Println("Nama pas")
	} else {
		fmt.Println("nama terlalu pendek")
	}
}
