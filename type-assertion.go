package main

import "fmt"

func random() interface{} {
	return "ups"
}

func main() {
	//type assertion merupakan kemampuan tipe data menjadi tipe data lain
	//sering digunakan ketika bertemu interface kosong

	var result interface{} = random()
	var resultString string = result.(string) //type assertion
	//harus benar karena jika tidak maka akan panic

	fmt.Println(resultString)

	//menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
