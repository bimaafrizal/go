package Web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

// dapat meload file secara otomatis menjadi static file
// FileServer adalah Handle sehingga bisa ditambahkan kedalam server mux

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	//mengahpus prefix static agar bisa membuka file
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//mux.Handle("/", fileServer)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// golang Embed
// static file akan diembed ke dalam binary

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	//fileServer := http.FileServer(http.FS(resources))
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// file tidak dapat dibuka karena "resources" masuk ke prefix golang embed maka gunakan fs.Sub
}

// serveFile adalah fungsi untuk membuka file secara langsung
func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// golang embed untuk serve file

//go:embed resources/ok.html
var resourcesOK string

//go:embed resources/notfound.html
var resourcesNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourcesOK)
	} else {
		fmt.Fprint(writer, resourcesNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
