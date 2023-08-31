package main

import "fmt"

func main() {
	//tipe data integer dibagi menjadi dua yaitu uin sampai negatif, tipe data uint dari 0 sampai positif
	//tipe data float digunakan untuk menampung bilangan desimal, dibagi menjadi dua yaitu float, dan complex. complex hanya untuk matematika yang rumit

	//alias
	//byte = uint8
	//rune = int32
	//int = minimal int32 bit
	//uint = minimal uint32 bit
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)

	//boolean disebut bool
	fmt.Println("benar = ", true)
	fmt.Println("salah = ", false)

	//string tidak terbatas
	//kata kunci string direpresentasikan dengan kata kuncin string
	//nilai data string diapit dengan tanda petik dua
	fmt.Println("nama saya = ", "Bima Afrizal")

	//function untuk string
	fmt.Println(len("Bima Afrizal")) //menghitung jumlah karakter
	fmt.Println("Bima Afrizal"[0])   //mengambil karakter berdasarkan index, mengembalikan byte huruf dari index

	//variable
	// variabel hanya bisa menyimpan tipe data yang sama
	var name string = "Bima Afrizal"
	fmt.Println(name)
	// name = 2  //error karena tipe data berbeda

	var age = 23 //jika tidak deklarasi tipe data maka akan otomatis menyesuaikan tipe data, jika diisi dengan integer maka akan menjadi int, jika ingin merubah menjadi int8 maka harus di deklarasi
	fmt.Println(age)

	//kata kunci var tidak wajib digunakan
	country := "Indonesia"
	fmt.Println(country)

	//deklarasi multiple variable
	var (
		firstName = "Bima 2"
		lastName  = "Afrizal 2"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//constant => variabel yang nilainya tidak bisa diubah
	//gunakan const
	//wajib inisiasi datanya

	const firstName3 string = "Bima 3"
	const lastName3 = "Afrizal 3"
	const age3 = 23
	//tidak dipakai tidak ada error

	//multiple constant
	const (
		firstName4 string = "Bima 4"
		lastName4         = "Afrizal 4"
	)

	//konversi tipe data
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) //nilai8 akan error karena melebihi batas
	//jika melebihi batas maka akan kembali ke nilai paling bawah(bisa negatif)

	var name5 = "Bima"
	var e = name5[0] 
	//mengambil byte dari string
	var eString = string(e)

	fmt.Println(eString)

	//type declaration
	//membuat alias dari tipe data yang sudah ada
	type NoKTP string //membuat alias dari string dengan nama NoKTP
	type Married bool //membuat alias dari bool dengan nama Married

	var noKtpBima NoKTP = "123456789"
	var marriedStatus Married = false
	fmt.Println(noKtpBima)
	fmt.Println(marriedStatus)

	//operasi matematika
	a := 10
	b := 20
	c := a + b
	fmt.Println(c)
	//augmented assignment
	a += 10 //a = a + 10
	fmt.Println(a)
	
	//unary operator
	i := 0
	i++
	fmt.Println(i)

	//comparison operator
	name6 := "Bima"
	name7 := "Bima"
	fmt.Println(name6 == name7) //true

	//operasi boolean
	var nilaiAkhir = 80
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	fmt.Println(lulusNilaiAkhir && lulusAbsensi) //true
}
