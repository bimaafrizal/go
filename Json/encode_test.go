package Json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// melakukan konversi data ke JSON menggunakan functiom json>Marshal(interface{

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Bima")
	logJson(1)
	logJson(true)
	logJson([]string{"Bima", "Afrizal", "Malna"})
}

// contoh diatas tidak sesuai karena harusnta adalah array atau object
