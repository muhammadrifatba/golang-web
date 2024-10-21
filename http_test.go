package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func TestHt(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello",nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString :=string(body)
	fmt.Println(bodyString)
}