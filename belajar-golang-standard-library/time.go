package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)

	utc := time.Date(1999, time.May, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "1999-05-01 12:00:00"

	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}
}
