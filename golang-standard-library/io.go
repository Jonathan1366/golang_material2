// IO atau singkatan dari Input Output, merupakan fitur di Golang yang digunakan sebagai standard untuk proses Input Output
// Di Golang, semua mekanisme input output pasti mengikuti standard package io
// https://pkg.go.dev/io

package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	// short name, 1 words
	var nama string
	fmt.Println("masukkan nama anda")
	fmt.Scanln(&nama)
	fmt.Println("jadi nama anda adalah", nama)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("masukkan nama panjang anda: ")
	longName,_:= reader.ReadString('\n')
	fmt.Println("jadi nama panjang saya adalah", longName)
}