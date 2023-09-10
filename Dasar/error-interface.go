package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai1 int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai1 / pembagi, nil //nil sebagai pengganti error
	}
}

func main() {
	//golang memiliki interdace yang digunakan sebagai kontrak untuk membuat sebuah error

	//biasa dilakukan dengan cara seperti ini
	hasil, err := pembagian(100, 2)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
