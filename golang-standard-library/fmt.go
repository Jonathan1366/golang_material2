// Sebelumnya kita sudah sering menggunakan package fmt dengan menggunakan function Println
// Selain Println, masih banyak function yang terdapat di package fmt, contohnya banyak digunakan untuk melakukan format
// https://pkg.go.dev/fmt

package main

import "fmt"

func main() {

	firstname:="Jonathan"
	lastname:="Farrel Emanuel"

	fmt.Printf("My name is %s %s", firstname, lastname)

	
}