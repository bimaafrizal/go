package Json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// kadang json berasal dari io.Reader
// kita bisa membaca data dari stream langsung di json
// untuk membuat json decoder dari stream, kita bisa menggunakan json.NewDecoder(io.Reader)
// untuk membaca isi inputan, kita bisa menggunakan decoder.Decode(interface{})

func TestStreamingDecode(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	_ = decoder.Decode(customer)
	fmt.Println(customer)
}

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("customer2.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Bima",
		MiddleName: "Afrizal",
		LastName:   "Malna",
		Age:        30,
		Married:    true,
	}

	_ = encoder.Encode(customer)
}
