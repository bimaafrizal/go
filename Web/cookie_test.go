package Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// http merupakan stateless artinya server tidak menyimpan informasi tentang client agar lebih mudah melakukan scalability di sisi server
// agar server mengingat client maka digunakan cookie
// cookie merupakan data yang disimpan di sisi client
// requset selanjutnya client akan membawa cookie secara otomatis
// cookie dibuat di server dan dikirim ke client
// cookie boleh lebih dari satu tapi jangan terlalu banyak karena akan membuat lama

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	// nama cookie
	cookie.Name = "bima-cookie"
	// value cookie
	cookie.Value = request.URL.Query().Get("name")
	// cookie diakses dimana saja, jika '/' maka bisa diakses di semua halaman
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Cookie created")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("bima-cookie")
	if err != nil {
		fmt.Fprint(writer, "No cookie found")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// unit test cookie
func TestSetCookie(t *testing.T) {
	requset := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookie?name=bima", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, requset)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	requset := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get-cookie", nil)
	cookie := new(http.Cookie)
	cookie.Name = "bima-cookie"
	cookie.Value = "bima"
	requset.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, requset)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
