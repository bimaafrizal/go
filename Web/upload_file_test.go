package Web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// multipart => form yang bisa mengirimkan file

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	templatesCache.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	//request dan max memory jika dikomen maka akan diberikan nilai defaultnnya
	//request.ParseMultipartForm(10 << 20) // 10mb
	// memiliki 3 return yaitu filenya, headernya dan errornya
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	//defer file.Close()
	//fmt.Println("File Name:", fileHeader.Filename)
	//fmt.Println("File Size:", fileHeader.Size)
	//fmt.Println("File Header:", fileHeader.Header)
	//folder menyimpan file
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	//defer fileDestination.Close()
	// pindahkan file ke folder
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	//mengambil bukan file
	name := request.PostFormValue("name")
	templatesCache.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/tes.png
var uploadFileTest []byte

func TestUploadFormClient(t *testing.T) {
	//inisiasi body dalam bentuk byte
	body := new(bytes.Buffer)
	// merubah bentuk body dalam bentuk multipart
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Bima")
	// mengambil form file
	file, _ := writer.CreateFormFile("file", "contoh.png")
	// mengambil file yang akan diupload
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	// harus set content type agar tidak gagal upload (seprti multipart/form-data di html)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
