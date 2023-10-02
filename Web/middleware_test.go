package Web

import (
	"fmt"
	"net/http"
	"testing"
)

// fitur dimana sebelum dan setelah sebuah handler dieksekusi, kita bisa menambahkan fungsi yang akan dieksekusi
// di golang tidak ada middleware ttapi kita bisa membuatnya sendiri menggunakan handler
type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// sebelum handler dieksekusi
	println("Before Handler")
	middleware.Handler.ServeHTTP(writer, request)
	// setelah handler dieksekusi
	println("After Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		panic("Ups")
	})

	//kebutuhan dari handler adalah pointer sehingga harus menambahkan &
	logMiddleware := &LogMiddleware{Handler: mux}
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
