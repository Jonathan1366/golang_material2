package golangjson

import (
	"encoding/json"
	"fmt"

	// "strings"
	"testing"
)

func LogJson(data interface{})  {
	bytes,err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	LogJson("Jonathan")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Jonathan", "Farrel","Emanuel"})
}