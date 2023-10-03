package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HttpRouter hanya untuk router
// tambahkan httprouter dan testify untuk testing

// sebuah struct router yang merupakan implementasi dari http.Handler
// untuk membuat router kita bisa menggunakan function httprouter.New() yang akan mengembalikan router pointer
func main() {
	router := httprouter.New()
	// kita bisa menentukan http method
	// perbedaan httpHandler dengan httprouter adalah terpat params
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Httprouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
