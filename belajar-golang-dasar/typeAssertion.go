package main

import "fmt"

func random() any {
	return true
}

func main() {
	data := random()

	// dataString := data.(string)

	// fmt.Println(dataString)

	switch value := data.(type) {
	case string:
		fmt.Println(data.(string))
	case int:
		fmt.Println(data.(int))
	default:
		fmt.Println("Uknown", value)
	}
}
