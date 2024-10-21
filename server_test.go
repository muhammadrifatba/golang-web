package golangweb

import (
 "testing"
 "net/http"
)
func TestServe(t *testing.T){
	server := http.Server{
		Addr:"localhost:8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}