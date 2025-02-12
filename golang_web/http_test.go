package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func HelloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello world")
	
}

func TestHelloWorld(t *testing.T)  {
	
	request:=httptest.NewRequest("GET","http://localhost:8080/hello", nil)
	recorder:= httptest.NewRecorder()

	HelloWorld(recorder, request)

	response:=recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyStr:=string(body)
	fmt.Println(bodyStr)
}