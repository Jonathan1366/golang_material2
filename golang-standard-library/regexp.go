// Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
// Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
// https://github.com/google/re2/wiki/Syntax
// https://golang.org/pkg/regexp/
//search filter product

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {


	//daftar produk
	products:= []string{"NIKE AiR 123","Zenbook","Macbook Air M2","Iphone 15 ProMax"}
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Println("Cari produk anda:")

	// membaca inptu dari pengguna
	scanner.Scan()
	input:=scanner.Text()

	pattern:= `(?i)^[A-Za-z0-9\s]+$`

	// KOMPILASI POLA REGEX UNTUK PENCARIAN PRODUK HURUF & ANGKA
	regex, err:= regexp.Compile(pattern)
	if err!= nil{

		fmt.Println("Error compiling regex:",err)
		return
	}

	// Pencarian produk yang cocok dengan input pengguna
	fmt.Println("Hasil:")
	for _, produk:= range products{
		if regex.MatchString(produk) && strings.Contains(strings.ToLower(produk),strings.ToLower(input)) {
			fmt.Println("Result:", produk)
		}
	}


}
