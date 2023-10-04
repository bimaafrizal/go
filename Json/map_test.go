package Json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// untuk atribute json yang dinamis kita bisa menggunakan map[string]interface{}, atribut merupakan key
// tidak mendukung json tag

func TestJSONMap(t *testing.T) {
	jsonRequest := `{"FirstName":"Bima","MidleName":"Afrizal","LastName":"Malna","Age":30,"Married":true,"Hobbies":["Reading","Writing","Coding"],"Address":[{"Street":"Jalan Raya","City":"Jakarta","Province":"DKI Jakarta"},{"Street":"Jalan Raya","City":"Bandung","Province":"Jawa Barat"}]}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["FirstName"])
	fmt.Println(result["Address"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Indomie",
		"price": 2000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
