package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Local())

	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formater:= "2006-01-02 15:04:05"
	value:= "2024-02-03 14:05:06"
	valueTime, err := time.Parse(formater, value)
	if  err!=nil {
		fmt.Println("Error", err.Error())
		}else{
			fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Second())



}