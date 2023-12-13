package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(4.4))
	fmt.Println(math.Floor(5.5))

	fmt.Println(math.Round(4.6))

	fmt.Println(math.Max(10, 15))

	fmt.Println(min(10, 5))
}
