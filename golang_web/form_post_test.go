package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//manual
func FormPost(w http.ResponseWriter, r *http.Request) {
	err:=r.ParseForm()
	if err != nil {
		panic(err)
	}
	first_name := r.PostForm.Get("first_name")
	last_name := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s",first_name, last_name)
} 

func TestFormPost(t *testing.T)  {
	requestBody:=strings.NewReader("first_name=Jonathan&last_name=Farrel")
	request:=httptest.NewRequest(http.MethodPost,"http://localhost:8080",requestBody)
	request.Header.Add("content-type","application/x-www-form-urlencoded")
	recorder:=httptest.NewRecorder()

	FormPost(recorder,request)

	response:=recorder.Result()
	body, _:= io.ReadAll(response.Body)
	fmt.Println(string(body))

}