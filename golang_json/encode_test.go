package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	w, _ := os.Create("customer.json")
	encoder:=json.NewEncoder(w)

	customer:=CreditCard{
		Number: "123123",
		CVV: "12344",
		Issuer: "benar",
	}
	encoder.Encode(customer)
	fmt.Println(customer)
}