package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "jonathan, farrel, emanuel\n" +
		"hai, kamu,siapa\n" +
		"anda, dia, halo"

	reader := csv.NewReader(strings.NewReader(csvString))

	for{
		record, err:= reader.Read()
		if err==io.EOF {
			break
		}
		fmt.Println(record)
	}

}