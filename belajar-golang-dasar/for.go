package main

import "fmt"

func main() {
	// counter := 1

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Perulangan ke:", counter)
	// 	counter++
	// }

	// fmt.Println("Selesai")

	names := []string{"john", "Wilkes", "Booth"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for _, name := range names {
		fmt.Println(name)
	}
}
