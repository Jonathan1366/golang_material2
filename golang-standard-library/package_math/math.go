// math.Round(float64)
// Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat

// math.Floor(float64)
// Membulatkan float64 kebawah

// math.Ceil(float64)
// Membulatkan float64 keatas

// math.Max(float64, float64)
// Mengembalikan nilai float64 paling besar

// math.Min(float64, float64)
// Mengembalikan nilai float64 paling kecil
package main
import ("fmt"
"math")

func main() {
	fmt.Println(math.Ceil(2.46))
	fmt.Println(math.Floor(2.46))
	fmt.Println(math.Round(2.46))
	fmt.Println(math.Max(10,46))
	fmt.Println(math.Min(10,46))

}