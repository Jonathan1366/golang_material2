// Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
// Kita bisa menggunakan errors.Is() untuk mengecek jenis type error nya

package main

import (
	"errors"
	"fmt"
)


var(ValidationError= errors.New("Empty")
NotFoundError=errors.New("not found error")
)

func GetById(id string) error {
	if id==""{
		return ValidationError
	} 
	if id!="jonathan"{
		return NotFoundError
	}
	return nil 
}

func main() {
	var userID string
	fmt.Print("Enter user id: ")
	_, err:= fmt.Scanln(&userID)
	if err!=nil {
		fmt.Println("error reading input:",err)
		return
	}
	err= GetById(userID)
	if err!=nil {
		if errors.Is(err,ValidationError) {
			fmt.Println("Validation error:",err)
		} else if errors.Is(err,NotFoundError){
			fmt.Println("Not found error:", err)
		} else{
			fmt.Println("Error",err)
		}
		return 
	}
	fmt.Println("user id is valid !")


}