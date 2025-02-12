// package main

// import (
// 	"bufio"
// 	"fmt"
// 	helper "golang_software_testing/golang-sql-injection-usecase" // Sesuaikan dengan path yang benar
// 	"os"
// 	"strings"
// )

// // Anda perlu menyesuaikan path di atas dengan struktur direktori dan nama file yang benar.

// func main() {
//     fmt.Print("Login User\n")

//     fmt.Print("Masukkaan email: ")
//     email := getUserInput()

//     // Minta pengguna untuk memasukkan password
//     fmt.Print("Masukkan password: ")
//     password := getUserInput()

//     //validasi login
//     err :=helper.ValidateLogin(email, password) // Menggunakan nama file login_user sebagai package
//     if err != nil {
//         fmt.Println("Login gagal", err)
//     } else {
//         fmt.Println("Login success untuk email:", email)
//     }
// }

// func getUserInput() string {
//     reader := bufio.NewReader(os.Stdin)
//     input, err := reader.ReadString('\n')
//     if err != nil {
//         panic(err)
//     }

//     // hilangkan karakter newline dari input
//     return strings.TrimSpace(input)
// }
 