package main

import "fmt"

type ipk interface{
	rata()float64
}

type hasil struct{
	latihan float64
	uts float64
	uas float64
}

func (h hasil) rata() float64 {
	return (h.latihan*30/100 + h.uas*40/100 + h.uts*30/100) /3
}

func main() {
	var latihan, uts, uas float64

	fmt.Println("Masukkan nilai anda\n")
	fmt.Scanln(&latihan, &uts, &uas )

	h:= hasil {latihan, uts, uas}
	nilaiIPK := h.rata()
	fmt.Printf("Nilai ipk anda adalah %.2f\n", nilaiIPK)
	

	
}

