package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name:=r.URL.Query().Get("name")
	if name=="" {
		//http.badrequest
		w.WriteHeader(400)
		fmt.Fprint(w,"name is empty")
	}else{
		w.WriteHeader(200)
		fmt.Fprintf(w,"Hi %s", name)
	}
}
func TestResponseCodeInvalid(t *testing.T){
	request:=httptest.NewRequest("GET","http://localhost:8080", nil)
	recorder:=httptest.NewRecorder()

	ResponseCode(recorder, request)

	response:=recorder.Result()
	body,_:=io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T){
	request:=httptest.NewRequest("GET","http://localhost:8080/?name=Jonathan", nil)
	recorder:=httptest.NewRecorder()

	ResponseCode(recorder, request)

	response:=recorder.Result()
	body,_:=io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	fmt.Println(string(body))
}