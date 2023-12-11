package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Error pembagi 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, error := pembagian(100, 0)

	if error == nil {
		fmt.Println("Hasil =", hasil)
	} else {
		fmt.Println("Error ", error)
	}
}
