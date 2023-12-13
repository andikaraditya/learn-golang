package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("eko khan", "eko"))
	fmt.Println(strings.Split("Eko Khan", " "))
	fmt.Println(strings.ToLower("EKO"))
	fmt.Println(strings.ToUpper("eko"))
	fmt.Println(strings.Trim("      Eko       ", " "))
	fmt.Println(strings.ReplaceAll("eko eko eko eko", "eko", "budi"))
}
