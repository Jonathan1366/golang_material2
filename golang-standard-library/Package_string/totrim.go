package main

// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.

import (
	"fmt"
	
	"strings"
)


func main() {

	fmt.Println(strings.Trim("    hai    ", " "))
	fmt.Println(strings.ReplaceAll("Jonathan Farrel Emanuel", "Jonathan","NATHAN"))

	
}