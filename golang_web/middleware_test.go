package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("before execute handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("after execute handler")
}

 type ErrorHandler struct{
	Handler http.Handler
 }

 func(errorhandler ErrorHandler) ServeHTTP(w http.ResponseWriter, r*http.Request){
	defer func(){
		err:=recover()
		if err != nil {
			fmt.Println("terjadi error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s",err) 
		}
	}()
	errorhandler.Handler.ServeHTTP(w, r)
 }

func TestMiddleWare(t*testing.T)  {
	mux:=http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(w, "hello middleware")

	})
	mux.HandleFunc("/foo",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("foo executed")
		fmt.Fprint(w, "hello foo")

	})
	mux.HandleFunc("/panic",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("foo executed")
		panic("upps")

	})
	
	LogMiddleware:= &LogMiddleware{
		Handler: mux,
	}

	errorhandler:=ErrorHandler{
		Handler: LogMiddleware,
	}

	server:=http.Server{
		Addr: "localhost:8080",
		Handler:errorhandler ,
	}

	err:= server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}