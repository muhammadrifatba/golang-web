package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request){
	name:=request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprintf(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T){
	request := httptest.NewRequest("GET", "http:/localhost:8080/hello",nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result
	body,_:= io.ReadAll(response().Body)
	bodyString := string(body)
	fmt.Printf(bodyString)
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request){
	firstname:=request.URL.Query().Get("first_name")
	lastname:=request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
		
}

func TestMultipleQueryParameter(t *testing.T){
	request := httptest.NewRequest("GET", "http:/localhost:8080/hello?first_name=rifat&last_name=Bagus",nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)
	response := recorder.Result
	body,_:= io.ReadAll(response().Body)
	bodyString := string(body)
	fmt.Printf(bodyString)
}

func MultipleParameterValue(writer http.ResponseWriter, request *http.Request){
	query:=request.URL.Query()
	names := query["name"]
	
	fmt.Fprintf(writer, strings.Join(names," "))
		
}

func TestMultipleValueParameter(t *testing.T){
	request := httptest.NewRequest("GET", "http:/localhost:8080/hello?name=Rifat&name=Bagus&name=Adikusuma",nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)
	response := recorder.Result
	body,_:= io.ReadAll(response().Body)
	bodyString := string(body)
	fmt.Printf(bodyString)
}