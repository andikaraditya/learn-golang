package main

import "fmt"

func main() {
	// person := map[string]string{
	// 	"name":    "Eko",
	// 	"address": "Jakarta",
	// }

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])
	// fmt.Println(person["age"])

	book := make(map[string]string)

	book["title"] = "Harry Potter"
	book["author"] = "JK Rowling"
	book["wrong"] = "wrong key"
	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)

}
