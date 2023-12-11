package main

import (
	"belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("John")

	fmt.Println(result)
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) not working

	fmt.Println(database.GetDatabase())
}
