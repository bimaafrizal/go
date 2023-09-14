package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

// untuk membuat goutotine cukup dengan menambahkan keyword go sebelum nama function maka function tersebut akan berjalan async
// aplikasi akan berjalan ke kode program selanjutnya tanpa menunggu gotroutine selesai
// jika fungsinya mengembalikan nilai maka tidak bisa menggunakan goroutine karena gorotine tidak bisa menangkap valuenya
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("TestCreateGoroutine")

	//tambahkan sleep agar aplikasi tidak langsung berhenti sebelum goroutine selesai
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGorotine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

//hasilnya tidak error karena leptopnya multicore sehingga berjalan secara pararel dan concurren, tapi sukses dijalankan
