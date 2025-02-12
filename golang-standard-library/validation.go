// api

package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
	 Name string `required:"true" max:"10"`
}

type Person struct{
	Name 		string `required:"true" max:"10"`
	Email 	string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`

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

func IsValidation(value any) bool {
	 t:=reflect.TypeOf(value)
	 for i := 0; i < t.NumField(); i++ {
		f:=t.Field(i)
		if f.Tag.Get("required") == "true"{
			data:= reflect.ValueOf(value).Field(i).Interface()
			if data=="" {
				return false
			}
		}
	 }
	 return true
}

func main() {

	// var s struct{
	// 	a string
	// 	b int
	// 	c float64
	// }
	// t:=reflect.TypeOf(s)
	// println(t.NumField())
	
	ReadFile(Sample{"Jonathan"})
	ReadFile(Person{"Farrel","halo","satu"})

	person:= Person{
		Name: "Jonathan",
		Email: "",
		Address: "Jakarta",
	}
	fmt.Println(IsValidation(person))
}
