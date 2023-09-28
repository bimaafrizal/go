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

// handlerFunc hanya bisa satu endpoint
// jika ingin lebih dari satu endpoint maka harus menggunakan serveMux
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hi")
		if err != nil {
			panic(err)
		}
	})

	//url patt
	// ketika url diakhiri / maka semua url tersebut yang dibelakang / akan dijalankan
	// misal images/ kemudian diketikan images/bima maka akan dijalankan images/
	// jika ada url yang lebih panjang maka akan dijalankan yang lebih panjang
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Images")
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/images/tumbnail", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Images Tumbnail")
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
		fmt.Fprintln(writer, request.Host)
		fmt.Fprintln(writer, request.Header)
		fmt.Fprintln(writer, request.Body)
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
