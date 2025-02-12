// api

package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
	 Name string `required:"True" max:"10"`
}

type Person struct{
	Name string `required:"True" max:"10"`
	Email string `required:"True" max:"10"`
	Address string `required:"True" max:"10"`

}
func ReadFile(nilai any){
	TipeNilai :=reflect.TypeOf(nilai)
	// var TipeNilai reflect.Value = reflect.TypeOf(nilai)
	// fmt.Println("Type name", TipeNilai.Name())
	for i := 0; i < TipeNilai.NumField(); i++ {
		NilaiField:=TipeNilai.Field(i)
		fmt.Println(NilaiField.Name, "with type", NilaiField.Type)
		fmt.Println(NilaiField.Tag.Get("required"))
		fmt.Println(NilaiField.Tag.Get("max"))

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
