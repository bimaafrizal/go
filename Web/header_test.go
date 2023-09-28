package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// informasi tambahan yang dikirimkan oleh client dan server
// tidak key sensitive

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	writer.Write([]byte("content type : " + contentType))
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	// untuk menambahkan header
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// menambhakan header pada response
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("x-power-by", "bima")
	fmt.Fprintln(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	// untuk menambahkan header
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	powerdBy := response.Header.Get("x-power-by")
	fmt.Println(string(body))
	fmt.Println(powerdBy)
}
