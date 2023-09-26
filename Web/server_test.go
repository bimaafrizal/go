package Web

import (
	"net/http"
	"testing"
)

// server adalah struct yang terdapat di package net/http digunakan untuk representasi web server
// harus menentukan host dan portnya defaultnya localhost:8080
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// menjalankan web server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
