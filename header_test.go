package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request){
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func ResponseHeader (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Rifat Bagus")
	fmt.Fprint(writer, "Ok")
}


func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}


func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("first_name=Rifat&last_name=Bagus")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/",requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	
	recorder := httptest.NewRecorder()
	FormPost(recorder, request)
	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/",nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)
	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/",nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)
	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}