package Json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street   string
	City     string
	Province string
}

// json array dipresentasikan dalam bentuk slice
// konversi dari json atau ke json dilakunan secara otomatis oleh package json menggunakan tipe data slice
type Customer2 struct {
	FirstName string
	MidleName string
	LastName  string
	Age       int
	Married   bool
	Hobbies   []string
	Address   []Address
}

func TestJsonArray(t *testing.T) {
	customer := Customer2{
		FirstName: "Bima",
		MidleName: "Afrizal",
		LastName:  "Malna",
		Age:       30,
		Married:   true,
		Hobbies:   []string{"Reading", "Writing", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Bima","MidleName":"Afrizal","LastName":"Malna","Age":30,"Married":true,"Hobbies":["Reading","Writing","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer2{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

// json array complex
func TestJsonArrayComplex(t *testing.T) {
	customer := Customer2{
		FirstName: "Bima",
		MidleName: "Afrizal",
		LastName:  "Malna",
		Age:       30,
		Married:   true,
		Hobbies:   []string{"Reading", "Writing", "Coding"},
		Address: []Address{
			{
				Street:   "Jalan Raya",
				City:     "Jakarta",
				Province: "DKI Jakarta",
			},
			{
				Street:   "Jalan Raya",
				City:     "Bandung",
				Province: "Jawa Barat",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Bima","MidleName":"Afrizal","LastName":"Malna","Age":30,"Married":true,"Hobbies":["Reading","Writing","Coding"],"Address":[{"Street":"Jalan Raya","City":"Jakarta","Province":"DKI Jakarta"},{"Street":"Jalan Raya","City":"Bandung","Province":"Jawa Barat"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer2{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Address)
	fmt.Println(customer.Address[0].City)
}

// decode json array
// melakukan encode dan decode langsung JSON Arraynya
// encode dan decode json array bisa menggunakan tipe data slice
func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Raya","City":"Jakarta","Province":"DKI Jakarta"},{"Street":"Jalan Raya","City":"Bandung","Province":"Jawa Barat"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		return
	}
	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:   "Jalan Raya",
			City:     "Jakarta",
			Province: "DKI Jakarta",
		},
		{
			Street:   "Jalan Raya",
			City:     "Bandung",
			Province: "Jawa Barat",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
