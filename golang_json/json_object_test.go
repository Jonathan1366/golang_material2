package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)


type Address struct{
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age int
	Married bool
	Hobbies [] string
	Addresses []Address 
}

func TestJSONObject(t *testing.T) {
	customer:=Customer{
		FirstName: "Jonathan",
		MiddleName: "Farrel",
		LastName: "Emanuel",
		Age: 20,
		Married: false,
	}
	bytes, _:= json.Marshal(customer)
	fmt.Println(string(bytes))
}

// {"FirstName":"Jonathan","MiddleName":"Farrel","LastName":"Emanuel","Age":20,"Married":false}

func TestJSONDecode(t*testing.T)  {
	jsonRequest:=`{"FirstName":"Jonathan","MiddleName":"Farrel","LastName":"Emanuel","Age":20,"Married":false}`
	jsonBytes:=[]byte(jsonRequest)

	customer:= &Customer{}
	err:=json.Unmarshal(jsonBytes,customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)





}