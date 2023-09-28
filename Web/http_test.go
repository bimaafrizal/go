package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// http test
// menjalankan testing tanpa menjalankan aplikasi web
// menggunakan NewRequest(method, url, body)
// NewRecoder digunakan untuk menampung response hasil testing
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
