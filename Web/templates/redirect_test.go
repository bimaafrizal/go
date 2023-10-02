package templates

import (
	"fmt"
	"net/http"
	"testing"
)

// response code harus 3xx dan menambahkan header location
// tapi ada cara otomatis

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://bimaafrizal.github.io/", http.StatusTemporaryRedirect)
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	//logicnya
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
	// statusnya bisa temporary atau permanent
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
