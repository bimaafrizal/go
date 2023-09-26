package test

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// package Embed adalah fitur untuk mempermudah membaca isi files pada saat compile time secara otomatis dimasukkan isi filenya dalam variabel yang kita tuju
// import embed
// tambahkan //go:embed diikuti dengan nama files diatas variabele yang kita tuju
// varibael yang dituju akan secara otomatis akan berisi konten files yang kita inginkan secara otomatis
// tiak boleh disimpan dalam function

// embed string

//go:embed ..\version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

// embed files ke []byte
// dalam bahasa lain array
// cocok untuk melakukan embed files yang tidak bisa dibaca seperti gambar, video, music, dsbnya

//go:embed ..\uns1.jpeg
var uns1 []byte

func TestByte(t *testing.T) {
	// save data baru
	err := ioutil.WriteFile("uns_new.jpeg", uns1, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// jika ingin menjalankan kode dibawah pindahkan pada root dorectory kemudian comment code pada main.go
// embed tidak bisa membuka file yang berada di luar direktorinya

// embed multiple files
// tambahkan go:embed lebih dari satu baris
// tipe data varibelnya embed.FS

////go:embed ..\files\a.txt
////go:embed ..\files\b.txt
////go:embed ..\files\c.txt
//var files embed.FS
//
//func TestMultipleFiles(t *testing.T) {
//	a, _ := files.ReadFile("../files/a.txt")
//	fmt.Println(string(a))
//	b, _ := files.ReadFile("../files/b.txt")
//	fmt.Println(string(b))
//	c, _ := files.ReadFile("../files/c.txt")
//	fmt.Println(string(c))
//}

// path matcher
// untuk membaca multiple file dengan pattern yang sama

////go:embed ../files/*.txt
//var path embed.FS
//
//func TestPathMatcher(t *testing.T) {
//	dir, _ := path.ReadDir("../files")
//	for _, entry := range dir {
//		if !entry.IsDir() {
//			fmt.Println("Nama File : ", entry.Name())
//			//file, _ := path.ReadFile("../files/" + entry.Name())
//			//fmt.Println(string(file))
//		}
//	}
//}
