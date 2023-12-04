package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func logging() {
	fmt.Println("Logging")
}

func endApp() {
	fmt.Println("Application is ended")
	message := recover()

	fmt.Println("Error:", message)
}

func runApplication(error bool) {
	// defer logging()
	defer endApp()
	fmt.Println("Application is running")
	if error {
		panic("Panic Error")
	}
}

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))

	runApplication(true)
}
