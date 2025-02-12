package main

import (
	"kantin2/canteen/internal/db/driver"
)


func main() {
	driver.InitDBPool()
	defer driver.CloseDBPool()
}
