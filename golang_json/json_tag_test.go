package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)
//snake_case
type Product struct{
	Id string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"img_url "`
}

func TestJSONTag(t*testing.T)  {
	product:=Product{
		Id: "P001",
		Name: "Apple Mac Book Pro",
		ImageUrl: "https://example.com/img.png",

	}
	bytes, _:= json.Marshal(product)
	fmt.Println(string(bytes))
}


func TestJSONTagDecode(t *testing.T)  {
	jsonString:=`{"id":"P001","name":"Apple Mac Book Pro","img_url ":"https://example.com/img.png"}`
	jsonBytes:=[]byte(jsonString)

	product:=&Product{}
	err:=json.Unmarshal(jsonBytes,product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.ImageUrl)
	fmt.Println(product.Name)
}