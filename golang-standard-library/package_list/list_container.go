package main

import (
	"container/list"
	"fmt"
)

//double linkedlist
func main() {

	// manual
	var data *list.List= list.New()
	data.PushBack("Jonathan")
	data.PushBack("Farrel")
	data.PushBack("Emanuel")

	current:=data.Front()
	fmt.Println(current.Value)


	kedua:=current.Next()
	fmt.Println(kedua.Value)

	ketiga:=kedua.Next()
	fmt.Println(ketiga.Value)
	
	for e := data.Front(); e!=nil; e=e.Next() {
		fmt.Println(e.Value)
	}


}