package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//secara default golang menggunakan sistem passing data by value bukan reference
	//jika kita mengirimkan sebuah variabel ke dalam function, method, atau variabel lain sebenarnya yang dikirimkan adalah duplikasi dari variabel tersebut
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
	//hasilnya akan berbeda karena address2 merupakan duplikasi dari address1

	//pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
	//kita bisa gunakan & untuk membuat pointer

	address3 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address4 := &address3
	//bisa ditulis var address4 *Address = &address3

	address4.City = "Tasikmalaya"
	fmt.Println(address3)
	fmt.Println(address4) //ada bintang berarti pointer

	//operator *
	//merubah data refrence secara utuh
	//membuat variabel pointer baru dan merubah data referenccenya
	address5 := Address{"Pengandaran", "Jawa Barat", "Indonesia"}
	address6 := &address5
	var address7 *Address = &address5

	*address6 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println(address7)

	//membuat function new
	//function new hanya mengembalikan pointer ke data kosong, bukan data kosong itu sendiri
	alamat := new(Address)
	alamat.City = "Bandung"
	fmt.Println(alamat)

	//pointer di function
	//digunakan untuk membuat funtion yang bisa mengubah data asli
	var alamat2 = Address{"Subang", "Jawa Barat", ""}
	ChangeCountryToIndonesia(alamat2) //hasilnya akan tetap kosong karena yang dipassing adalah duplikasi dari alamat2
	fmt.Println(alamat2)

	//tambahna tanda &
	ChangeCountryToIndonesia2(&alamat2) //hasilnya akan berubah karena yang dipassing adalah pointer dari alamat2
	fmt.Println(alamat2)

	bima := Man{Name: "Bima"}
	bima.Married()

	fmt.Println(bima)
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

// gunakan * untuk merubah pointer
// untuk data struct yang besar gunakan pointer agar tidak membebani memory
func ChangeCountryToIndonesia2(address *Address) {
	address.Country = "Indonesia"
}

// sebenarnya data stuct yang diakses di dalam method adalah by value
// gunakanan pointer untuk menghemat pointer
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
