package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func getFullName() (string, string) {
	return "Barrack", "Obama"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khan"
	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func getGoodbye(name string) string {
	return "Goodbye " + name
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func filterName(name string) string {
	if name == "Dog" {
		return "..."
	} else {
		return name
	}
}

type BlackList func(string) bool

func registeruser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Please enter")
	}
}

func main() {
	sayHello()
	sayHello()
	sayHello()
	sayHelloTo("John", "Wilkes")
	fmt.Println(getHello("Jake"))

	// firstName, lastName := getFullName()

	// fmt.Println(firstName, lastName)

	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

	fmt.Println(sumAll(5, 5, 5, 5))

	numbers := []int{2, 2, 2, 2}

	fmt.Println(sumAll(numbers...))

	goodbye := getGoodbye

	fmt.Println(goodbye("Luke"))

	sayHelloWithFilter("Dog", filterName)

	blackList := func(name string) bool {
		return name == "Dog"
	}

	registeruser("John", blackList)

	registeruser("Dog", func(name string) bool {
		return name == "Dog"
	})
}
