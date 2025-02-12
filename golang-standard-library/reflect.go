// Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
// Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
// Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package reflec ini secara otodidak
// Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
// https://golang.org/pkg/reflect/

// package main

// import "fmt"

// type Pengguna struct{
// 	nama string
// 	umur int
// }

// type User []Pengguna

// func main() {
// 	var users User
// 	users= append(users, Pengguna{
// 		nama:"Alice", umur: 20})

// 		fmt.Println(users)

// }

package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
	 Name string
}

type Person struct{
	Name, email, address string
}
func ReadFile(nilai any){
	TipeNilai :=reflect.TypeOf(nilai)
	// var TipeNilai reflect.Value = reflect.TypeOf(nilai)
	// fmt.Println("Type name", TipeNilai.Name())
	for i := 0; i < TipeNilai.NumField(); i++ {
		NilaiField:=TipeNilai.Field(i)
		fmt.Println(NilaiField.Name, "with type", NilaiField.Type)
	}
}

func main() {

	var s struct{
		a string
		b int
		c float64
	}
	t:=reflect.TypeOf(s)
	println(t.NumField())
	
	ReadFile(Sample{"Jonathan"})
	ReadFile(Person{"Farrel","halo","satu"})
}
