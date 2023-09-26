package Web

import (
	"fmt"
	"net/http"
	"testing"
)

// handler untuk menerima hhtp request
// bentuknya interface, dimana di dalam kontraknya terdapat sebuah function bernama ServeHTTP() yang digunakan sebagai function yang akan diekseskusi ketika menerima http request
// HendlerFunc adalah function yang mengimplementasikan interface http.Handler
func TestHendler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
