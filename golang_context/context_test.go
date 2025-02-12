package golangcontext

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
		background:=context.Background()
		fmt.Println(background)
		todo:= context.TODO()
		fmt.Println(todo)
}

func TestContextWithValue(t*testing.T){
	contextA:= context.Background()
	
	contextB:= context.WithValue(contextA,"b","B")
	contextC:= context.WithValue(contextA,"c","C")
	
	contextD:= context.WithValue(contextB,"d","D")
	contextE:= context.WithValue(contextB,"e","E")
	
	contextF:= context.WithValue(contextC,"f","F")
	contextG:= context.WithValue(contextC,"g","G")

	fmt.Println(contextA) 
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("b"))





}


type User struct{
	ID int
	Name string
}


func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)  {
	// mendapatkan nilai user dari context
	user, ok:= ctx.Value("user").(User)
	if !ok{
		http.Error(w,"user not found", http.StatusInternalServerError)
		return
	}
	// tampilkan informasi user
	fmt.Fprintf(w,"Hello, %s!", user.Name)
}

func TestHandler(t*testing.T)  {
	// buat user dan context dengan user
	user:= User{ID:1, Name:"Jonathan"}
	ctx:=context.WithValue(context.Background(),"user", user)

	// buat request dan response recorder
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	// Panggil handler dengan context
	Handler(ctx, rec, req)

	// Periksa hasilnya
	result := rec.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200, got %d", result.StatusCode)
	}

	expected := "Hello, Jonathan!"
	if rec.Body.String() != expected {
			t.Fatalf("Expected body %q, got %q", expected, rec.Body.String())
	}
}


// package main

// import (
//     "context"
//     "fmt"
//     "log"
//     "net/http"
// )

// type TraceID string

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         // Simulasi pembuatan TraceID
//         traceID := TraceID("12345")

//         // Membuat context dengan TraceID
//         ctx := context.WithValue(r.Context(), "traceID", traceID)

//         // Meneruskan context ke handler berikutnya
//         logRequest(ctx, w, r)
//     })

//     http.ListenAndServe(":8080", nil)
// }

// func logRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
//     // Mendapatkan TraceID dari context
//     traceID, ok := ctx.Value("traceID").(TraceID)
//     if !ok {
//         http.Error(w, "Trace ID not found", http.StatusInternalServerError)
//         return
//     }

//     // Logging dengan TraceID
//     log.Printf("Received request with Trace ID: %s", traceID)
//     fmt.Fprintln(w, "Request logged.")
// }


// goroutine leaks
// func CreateCounter() chan int {
// 	destination:= make(chan int)
// 	go func() {
// 		defer close(destination)
// 		counter:=1
// 		for{
// 			destination<-counter
// 			counter++
// 		}
// 	}()
// 	return destination
// }

func CreateCounter(ctx context.Context) chan int {
	destination:= make(chan int)
	go func() {
		defer close(destination)
		counter:=1
		for{
			select{
			case <- ctx.Done():
				return
			default:
				destination<- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t*testing.T )  {
	fmt.Println("Total goroutine",runtime.NumGoroutine())
	
	parent:= context.Background()
	ctx, cancel:=context.WithCancel(parent)
	
	destination:=CreateCounter(ctx)
	for n:=range destination {
		fmt.Println("Counter",n)
		if n==10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context
	// time.Sleep(1*time.Second)
	// fmt.Println("Total goroutine", runtime.NumGoroutine())

	//ambil semua nilai dari channel hingga channel ditutup
	for range destination{
		//do nothing
	}
	time.Sleep(1*time.Second) // mengurangi waktu tidur untuk mempercepat eksekusi
	fmt.Println("Total goroutine",runtime.NumGoroutine())
}

// Penjelasan:
// Paket main: Kode berada di dalam paket main agar dapat dijalankan sebagai program Go yang mandiri.

// Fungsi CreateCounter:

// Membuat channel destination untuk mengirim nilai integer.
// Menjalankan goroutine yang mengirim nilai counter ke channel, dan berhenti ketika context ctx dibatalkan.
// Fungsi TestContextWithCancel:

// Mencetak jumlah goroutine yang berjalan sebelum dan sesudah eksekusi fungsi.
// Membuat context ctx dan fungsi pembatalan cancel dari context parent.
// Memulai goroutine dengan CreateCounter(ctx) dan mengeluarkan nilai dari channel destination.
// Memanggil cancel() setelah nilai counter mencapai 10 untuk menghentikan goroutine.
// Melakukan iterasi melalui sisa nilai dari channel destination untuk memastikan semua nilai diambil dan goroutine ditutup dengan benar.
// Mengurangi waktu tidur untuk mempercepat eksekusi dan kemudian mencetak jumlah goroutine yang berjalan.
// Fungsi main:

// Memanggil TestContextWithCancel untuk menjalankan tes.
// Dengan menambahkan loop tambahan setelah cancel() untuk memastikan semua nilai dari channel diambil, Anda memastikan bahwa goroutine di CreateCounter dihentikan dengan benar dan tidak ada kebocoran goroutine yang terjadi.
