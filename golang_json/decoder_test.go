package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Struct untuk koordinat lokasi
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Struct untuk lokasi
type Location struct {
	Street      string      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Coordinates Coordinates `json:"coordinates"`
}

// Struct untuk nama pengguna
type Name struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
}

// Struct untuk pekerjaan
type Job struct {
	Title      string `json:"title"`
	Descriptor string `json:"descriptor"`
	Area       string `json:"area"`
	Type       string `json:"type"`
	Company    string `json:"company"`
}

// Struct untuk kartu kredit
type CreditCard struct {
	Number string `json:"number"`
	CVV    string `json:"cvv"`
	Issuer string `json:"issuer"`
}

// Struct utama untuk setiap hasil
type Result struct {
	Message       string     `json:"message"`
	PhoneNumber   string     `json:"phoneNumber"`
	PhoneVariation string    `json:"phoneVariation"`
	Status        string     `json:"status"`
	Name          Name       `json:"name"`
	Username      string     `json:"username"`
	Password      string     `json:"password"`
	Emails        []string   `json:"emails"`
	Location      Location   `json:"location"`
	Website       string     `json:"website"`
	Domain        string     `json:"domain"`
	Job           Job        `json:"job"`
	CreditCard    CreditCard `json:"creditCard"`
	UUID          string     `json:"uuid"`
	ObjectID      string     `json:"objectId"`
}

type Data struct {
	Result []Result `json:"result"`
}


func TestStreamDecoder(t *testing.T) {
	reader, _:=os.Open("customer.json")
	decoder:=json.NewDecoder(reader)

	sample:= &Data{}
	decoder.Decode(sample)

	fmt.Println(sample)
}