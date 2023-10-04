package Json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json object di go-lang dipresentasikan dengan tipe data struct
// dimana tiap attribute di json object akan menjadi field di struct

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Bima",
		MiddleName: "Afrizal",
		LastName:   "Malna",
		Age:        30,
		Married:    true,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
