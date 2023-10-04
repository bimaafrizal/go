package Json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// mengembalikan json ke go-lang
// json.Unmarshal([]byte, interface{)
// byte[] adalah data json, sedangkan interface tempat menyimpan hsil konversi biasanya berupa pointer

func TestUnmarshal(t *testing.T) {
	jsonString := `{"FirstName":"Bima","MidleName":"Afrizal","LastName":"Malna","Age":30,"Married":true}`
	jsonBytes := []byte(jsonString)
	// gunakan pointer untuk menampung hasil konversi, jika tidak ditakutkan variabel customer tidak mendapatkan perubahan data
	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
}
