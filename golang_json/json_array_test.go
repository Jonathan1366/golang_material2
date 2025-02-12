package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	type Customer struct {
		FirstName  string
		MiddleName string
		LastName   string
		Age int
		Married bool
		Hobbies [] string
	}
	
	customer:= Customer {
		FirstName:"Jonathan",
		MiddleName:"Farrel",
		LastName:"Emanuel",
		Hobbies: []string{"Coding","Reading", "watch"},

	}
	bytes, _:=json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONDecodeArray(t *testing.T)  {
	jsonString:=`{"FirstName":"Jonathan","MiddleName":"Farrel","LastName":"Emanuel","Age":0,"Married":false,"Hobbies":["Coding","Reading","watch"]}`
	jsonBytes:=[]byte(jsonString)

	customer:=&Customer{}
	err:=json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)

}


func TestJSONArrayComplex(t*testing.T)  {
	customer:=Customer{
		FirstName: "Jonathan",
		Addresses: []Address{
			{
				Street:"Lodan" ,
				Country: "Indonesia",
				PostalCode: "8080",
			},
			{
				Street:"gading" ,
				Country: "Indonesia",
				PostalCode: "7070",
			},
			{
				Street:"Sudirman" ,
				Country: "Indonesia",
				PostalCode: "1000",
			},
		
		},
	}
	bytes, _:=json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString:=`{"FirstName":"Jonathan","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Lodan","Country":"Indonesia","PostalCode":"8080"},{"Street":"gading","Country":"Indonesia","PostalCode":"7070"},{"Street":"Sudirman","Country":"Indonesia","PostalCode":"1000"}]}`

	jsonBytes:=[]byte(jsonString)

	customer:=&Customer{}
	err:=json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString:=`[{"Street":"Lodan","Country":"Indonesia","PostalCode":"8080"},{"Street":"gading","Country":"Indonesia","PostalCode":"7070"},{"Street":"Sudirman","Country":"Indonesia","PostalCode":"1000"}]`

	jsonBytes:=[]byte(jsonString)

	addresses:=&[]Address{}
	err:=json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
	
}

func TestOnlyJSONArrayComplex(t*testing.T)  {
	addresses:=[]Address{
		{
			Street:"Lodan" ,
			Country: "Indonesia",
			PostalCode: "8080",
		},
		{
			Street:"gading",
			Country: "Indonesia",
			PostalCode: "7070",
		},
		{
			Street:"Sudirman" ,
			Country: "Indonesia",
			PostalCode: "1000",
		},
	
	}
	bytes, _:=json.Marshal(addresses)
	fmt.Println(string(bytes))
}



