package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//QUERY PARAMS
func SayHello(w http.ResponseWriter, r *http.Request) {
	name:=r.URL.Query().Get("nama")
	if name=="" {
		fmt.Fprint(w, "hello")
	}else {
		fmt.Fprintf(w, "hello %s", name)
	}
}

func TestQueryParam(t *testing.T)  {
	request:= httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?nama=Jonathan",nil)
	recorder:=httptest.NewRecorder()
	SayHello(recorder, request)

	response:=recorder.Result()
	body,_:=io.ReadAll(response.Body)

	fmt.Println(string(body))
}

//MULTIQUERY

func MultiQueryParam(w http.ResponseWriter, r *http.Request) {
	firstName:= r.URL.Query().Get("first_name")
	lastnName:= r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastnName)
}

func TestMultiQuery(t *testing.T) {
	request:= httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Jonathan&last_name=Farrel",nil)
	recorder:=httptest.NewRecorder()
	MultiQueryParam(recorder, request)

	response:=recorder.Result()
	body,_:=io.ReadAll(response.Body)

	fmt.Println(string(body))
}

//multiparamsvalues

func MultiParamValues(w http.ResponseWriter, r *http.Request)  {
	query:= r.URL.Query()
	names:=query["name"]
	fmt.Fprint(w,strings.Join(names," "))
}


func TestMultiParamValues(t *testing.T){
	request:= httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Jonathan&name=Farrel&name=Emanuel",nil)
	recorder:= httptest.NewRecorder()
	MultiParamValues(recorder, request)

	response:=recorder.Result()
	body,_:=io.ReadAll(response.Body)

	fmt.Println(string(body))
}