package main

import "fmt"

// import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func CreateMap(param ...interface{}) map[interface{}]interface{} {
	if len(param)%2 != 0 {
		return nil
	} else {
		var m map[interface{}]interface{} = make(map[interface{}]interface{})
		for i := 0; i < len(param); i += 2 {
			m[param[i]] = param[i+1]
		}
		return m
	}
}

func main() {
	//jika kita buat variabel dengan tipe data tertentu maka akan dibuatkan default valuenya
	//nil digunakan untuk data kosong
	//hanya bisa digunakan di beberapa tipe data seperti interdace, function, map, slice, pointer, channel
	var person map[string]string = nil
	fmt.Println(person)

	person2 := NewMap("")
	fmt.Println(person2)

	person3 := NewMap("Bima")
	fmt.Println(person3)

	imap := CreateMap("name", "Bima", "title", "Programmer", "age", 20)
	imap2 := CreateMap("name")
	fmt.Println(imap2)
	fmt.Println(imap)
}
