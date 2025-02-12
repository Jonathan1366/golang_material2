// Di Golang versi terbaru, terdapat fitur bernama Generic, fitur ini akan kita bahas khusus dikelas Golang Generic
// Fitur Generic ini membuat kita bisa membuat parameter dengan tipe yang bisa berubah-ubah, tanpa harus menggunakan interface kosong / any
// Salah satu package yang menggunakan fitur Generic ini adalah package slices
// Package slices ini digunakan untuk memanipulasi data di slice
// https://pkg.go.dev/slices

package main

import (
	"fmt"
	"slices"
)
func main() {
	names:= []string{"Jonathan", "farrel","emanuel"}
	values:= []int{100,90,80}
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names,"Jonathan"))
	fmt.Println(slices.Index(names,"farrel"))
	fmt.Println(slices.Contains(names,"emanuel"))
}