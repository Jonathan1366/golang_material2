package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{ "Jonathan","Farrel","emanuel" })
	_ = writer.Write([]string{ "Hans","Budi","luhut" })
	_ = writer.Write([]string{ "Thomas","alfa","edison" })

	writer.Flush()
}