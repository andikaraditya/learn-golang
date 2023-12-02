package main

import "fmt"

func main() {
	// months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// slice1 := months[2:4]
	// fmt.Println(slice1)
	// slice2 := months[6:]
	// fmt.Println(slice2)
	// slice3 := months[:6]
	// fmt.Println(slice3)
	// slice4 := months[:]
	// fmt.Println(slice4)

	// days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// fmt.Println(days)
	// daySlice1 := days[5:]
	// fmt.Println(daySlice1)

	// daySlice1[0] = "New Saturday"
	// daySlice1[1] = "New Sunday"
	// fmt.Println(days)
	// fmt.Println(daySlice1)

	// daySlice2 := append(daySlice1, "New Day")
	// daySlice2[0] = "Old Saturday"
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	// var newSlice []string = make([]string, 2, 5)

	// newSlice[0] = "Eko"
	// newSlice[1] = "Khan"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// fromSlice := days[:]
	// toSlice := make([]string, len(fromSlice), cap(fromSlice))

	// copy(toSlice, fromSlice)

	// fmt.Println(fromSlice)
	// fmt.Println(toSlice)

	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
