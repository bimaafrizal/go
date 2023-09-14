package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

// channel adalah tempat komunikasi antar goroutine secara synchronous
// untuk mendapatkan return value dari goroutine maka
// goruotine akan terblok saat mengirim data ke channel sampai data ada yang memanggil
// hanya bisa mengirim satu data
// channel hanya bisa menerima satu tipe data
// channel bisa diambil dari lebih dari satu goroutine tapi satu2
// channel bisa menerima dari lebih dari satu goroutine tapi satu2
// channel harus di close jika tidak digunakan -> gunakan close(channel)
// direpresentasikan dengan tipe data chan
// harus deklarasi tipe data terleebih dahulu
// mengirim data gunakan kode: channel <- data
// menerima data gunakan kode: data <- channel

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) //membnuat channel dengan tipe data string
	defer close(channel)         //menutup channel

	//channel <- "Hello" //mengirim data ke channel
	//data := <-channel //menerima data dari channel
	//fmt.Println(<- channel) //menerima data dari channel dan langsung di print

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Bima Afrizal" //mengirim data ke channel
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel //menerima data dari channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// channel sebagai parameter
// channel tidak perlu diganti menjadi reference
func GiveMeRespon(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Bima Afrizal" //mengirim data ke channel
}

func TestChannelAsPareameter(t *testing.T) {
	channel := make(chan string) //membuat channel dengan tipe data string
	defer close(channel)         //menutup channel

	go GiveMeRespon(channel)

	data := <-channel //menerima data dari channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// channel in dan out
// channel in -> hanya bisa menerima data
// channel out -> hanya bisa mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Bima Afrizal" //mengirim data ke channel
}
func OnlyOut(channel <-chan string) {
	data := <-channel //menerima data dari channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string) //membuat channel dengan tipe data string
	defer close(channel)         //menutup channel

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
