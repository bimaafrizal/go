package entity

type Comments struct {
	Id      int32
	Email   string
	Comment string
}

// repository pattern => mekanisme untuk mengeksapsulasin storage
// digunakan sebagai jembatan antara business logic dengan perintah SQL ke database
// model dalam golang dalam bentuk struct
// return lebih baik dikonversi terlebih dahulu ke struct Entity/Model sehingga kita tinggal menggunakan objectnya
