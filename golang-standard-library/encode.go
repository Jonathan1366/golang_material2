// Golang menyediakan package encoding untuk melakukan encode dan decode
// https://pkg.go.dev/encoding
// Golang menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah base64, csv dan json

package main

import (
	"encoding/base64"
	"fmt"
)

// decode base64
func main() {
	value:="Jonathan Farrel Emanuel"
	encoded:=base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)
	
	decoded, err:=base64.StdEncoding.DecodeString(encoded)
	if err!=nil {
		fmt.Println("Error",err.Error())
	}else{
	fmt.Println(string(decoded))		 
	}
	 
}