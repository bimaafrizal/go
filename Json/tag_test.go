package Json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// secara default attribut yanf di struct dan json akan di mapping sesuai dengan nama atribut yang sama
// kadang terdapat penulisan yang berbeda json dengan snake_case sedangkan di go-lang PascalCase
// kita bisa menambahkan tag refelection pada struct untuk mengubah nama attribut di json

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Aple",
		Price:    "RP500.000",
		ImageUrl: "http://example.com/image.png",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","NAME":"Aple","price":"RP500.000","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	// untuk unmarshal tidak case sensitive, akan tetapi image url tidak masuk karena memakai _ sehingga perlu menambahkan _
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
