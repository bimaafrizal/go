package Goroutines

import (
	"fmt"
	"strconv"
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

// buffered channel
// kadang ads kasus dimana pengirim lebih cepat dari penerima, maka otomatis pengirim akan ikut lambat
// buffered channel adalah channel yang bisa figunakan untuk menampung data antrian
// buffer capacity => jumlah data yang bisa ditampung dalam antrian
// jika buffer penuh maka data akan diminta menunggu
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) //membuat channel dengan tipe data string dan buffer 3
	defer close(channel)            //menutup channel

	go (func() {
		channel <- "Bima"
		channel <- "Afrizal"
		channel <- "Malnaa"
	})()

	go (func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	})()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

// range cahnnel
// untuk kasus dimana kita mengirim data secara terus menerus ke channel dan tidak tau kapan berhenti
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1000; i < 100000; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	//perulangannya akan berhenti ketika channelnya di close
	for data := range channel {
		fmt.Println("menerima data", data)
	}

	fmt.Println("selesai")
}

// select channel
// mengambil data dari channel sekaligus
// digunakan untuk mendapatkan data tercepat dari beberapa channel yang berjalan bersamaan, jika bersamaan maka akan dipilih secara random
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespon(channel1)
	go GiveMeRespon(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		// default select channel
		// digunakan untuk menghindari deadlock karena data channel tidak ada
		// akan dieksekusi jika tidak ada data di channel
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
