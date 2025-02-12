package main

import (
	"fmt"
	"strconv"
)

func main() {
	v:= "10"
	if s, err:= strconv.Atoi(v); err==nil{
		fmt.Printf(" %T , %v" , s, s)
	}

	hah:=true
	h:=strconv.FormatBool(hah )
	fmt.Printf(" %T, %v \n", h,h )



	jonathan, err:=strconv.ParseBool("true")
	if err==nil {
		fmt.Println(jonathan)
	}else {
		fmt.Println("Error",err.Error())
	}

	//convert binary
	binary:= strconv.FormatInt(48, 2)
	fmt.Println(binary)

	// convert ke string
	binary2:=strconv.Itoa(999)
	fmt.Println(binary2)

}

