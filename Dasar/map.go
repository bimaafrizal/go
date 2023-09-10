package main

import "fmt"

func main() {
	//key value
	//key bersifat unique
	//tidak ada batasan jumlah data, jika kita menggunakan kata kunci yang sama maka akan mengganti data sebelumnya
	person := map[string]string{
		"name":    "Bima",
		"address": "Solo",
	}

	//menambah / merubah data
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])

	//fungsi map
	fmt.Println(len(person))  //menghitung jumlah data
	delete(person, "address") //menghapus data
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	//book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Bima"
	fmt.Println(book)
}
