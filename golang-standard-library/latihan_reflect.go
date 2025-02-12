// Soal: Dalam konteks Golang, reflect adalah paket yang memungkinkan kita untuk melakukan introspeksi dan manipulasi pada objek pada waktu eksekusi. Jelaskan bagaimana Anda dapat menggunakan paket reflect untuk mendapatkan informasi tentang variabel yang tipe datanya tidak diketahui pada waktu kompilasi. Berikan contoh kode yang menunjukkan bagaimana Anda dapat menggunakan reflect untuk:

// Menentukan tipe dari variabel.
// Mengakses nilai dari variabel tersebut.
// Memodifikasi nilai dari variabel tersebut jika itu adalah pointer.

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x int = 40
	var px *int = &x // buat pointer ke x

	//1. menentukan tipe dari variable
	
	t:=reflect.TypeOf(x)
	fmt.Println("Tipe:",t)

	//2. Mengakses nilai dari variabel tersebut.
	v:=reflect.ValueOf(x)
	fmt.Println("Value: ",v)

	//3 modifikasi nilai dari variable tersebut jika itu adalah pointer
	
	vp:= reflect.ValueOf(px)
	if vp.Kind()== reflect.Ptr {
		// mendapatkan nilai yang pointer tunjuk
		e:= vp.Elem()
		if e.CanSet() {
			// modifikasi nilai pointer
			e.SetInt(7)
		}
	}
	fmt.Println("Nilai baru dari x:",x)

}