package Web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// xss adalah salah satu security issie diamana orang bisa dengan sengaja memasukan parameter yang mengandung js agar dirender oleh halaman kita
// biasanya XSS mencuri cookie browser pengguna yang sedang mengakses website kita
// auto escape adalah fitur yang dimiliki go untuk menghindari XSS, go dapat mendetaksi data yang ditampilkan di template, jika mengandung tag2 html atau script maka akan diescape secara otomatis
/// gunakan html template jangan gunakan text template

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	templatesCache.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Golang Auto Escape",
		//tagnya akan di escape
		"Body": "<p>Ini adalah paragraf</p><script>alert('Anda di hack')</script>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// jika ingin menampilkan data html atau script atau css maka gunakan html template
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	templatesCache.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Golang Auto Escape",
		//tagnya akan di escape
		"Body": template.HTML("<p>Ini adalah paragraf</p><script>alert('Anda di hack')</script>"),
	})
}

func TestTemplateAutoEscapeServerDisabled(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// pastikan kita benar2 percaya terhadap sumber data yan gkita matikan auto espcapenya
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	templatesCache.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Golang Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
