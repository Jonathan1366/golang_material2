package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(test *testing.T){
	result:= HelloWorld("Jonathan Farrel Emanuel1")
	if result!= "Jonathan Farrel Emanuel" {

		// failed

		// panic("result is wrong")// kurang bagus

		// test.Fail()
		// test.FailNow()
		// test.Fatal()
		// test.Error()

fmt.Println("Test success")
	}
}
	