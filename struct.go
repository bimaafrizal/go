package main

import "fmt"

//struct
//struct adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//struct biasanya representasi data dalam program aplikasi yang kita buat
//data di struct disimpan dalam field
//mirip array tapi bisa beda tipe data

//gunakan uppercase untuk nama field dana nama struct
type Customer struct {
	Name, Address string
	Age           int
}

//function di struct
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var bima Customer
	bima.Name = "Bima Afrizal"
	bima.Address = "Indonesia"
	bima.Age = 22

	fmt.Println(bima)
	fmt.Println(bima.Name)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     22,
	}

	//tidak direkomendasikan karena urutan harus sesuai & jika ada field baru maka akan terjadi error
	budi := Customer{"Budi", "Indonesia", 22}

	fmt.Println(joko)
	fmt.Println(budi)

	bima.sayHi("Bima")
}
