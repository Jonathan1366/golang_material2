package main

import (
	"bufio"
	"fmt"
	"golang_software_testing/latihan_baru"
	"os"
)

func main() {
	// inisialisasi keywordstore dengan beberapa keyword
	keywords:=[]string {"golang","blockchain", "AI","drone","delivery"}
	keywordStore:=latihan_baru.NewKeyWordStore(keywords)

	//buat sebuah scanner untuk baca user input
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a keyword to search:")

	for scanner.Scan(){
		keyword:=scanner.Text()
		result, err:=keywordStore.SearchItem(keyword)
		if err!=nil {
			fmt.Println(err)
		}else {
			fmt.Println("Found:", result)
			keywordStore.AddToChart(result)
			fmt.Println("item added to cart")
		}
		fmt.Println("Enter another keyword to search or press CTRL+C to exit:")
	}
	// tampilkan isi cart sebelum keluar
	fmt.Println("Item in your cart:")
	for _, item:=	range keywordStore.ViewCart(){
		fmt.Println(item)
	}
}