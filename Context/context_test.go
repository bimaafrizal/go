package Context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan deadline
// context dibuat per request
// context untuk mempermudah kita meneruskan value dan sinyal antar proses
// wajib digunakan dan selalu dikirimkan setiap function dikirimkan
// cara kerja
// saat ada proses lain yang berjalan, contex akan mengirimkan sinyal cancel atau timeout ke proses lain maka proses lain akan menghentikan
// context.Background() : membuat contex kosong, tidak pernah timeout, dan tidak pernah dibatalkan biasanya digunakan pada main function atau dalam test atau ketika awal proses request terjadi
// context.TODO() : sama seperti Background() namun biasanya digunakan pada function yang belum jelas contexnya
func TestContext(t *testing.T) {
	//dalam membuat web tidak perlu membuat context secara manual karena sudah disediakan oleh golang
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// parent and child context
// parent context adalah context yang dibuat terlebih dahulu
// child context adalah context yang dibuat dari parent context, hanya bisa satu parent
// mirip pewarisan di oop
// akan selalu terhubung antara parent dan child
// jika kita melakukan pembatalan context A, maka semua child dan sub child context akan ikut dibatalkan
// yang dibatalkan hanya context yang bersangkutan dan turunannya, parentnya tidak di cancel
// begitu juga untuk mengirim data value, hanya bisa dari parent ke child
// immutable, setelah dibuat tidak bisa diubah
// ketika menambahkan data maka akan membuat child baru

// context with value
// value dengan key dan value
func TestContexWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// mengambil value
	// kalau tidak dapat akan mengambil parentnya
	// jika diparent tidak ada maka akan mengembalikan nil
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

// context with cancel
// digunakan ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses lain tersebut
// menggunakan goroutine lain untuk mengirimkan cancel
// akan membuat context baru
// untuk mengatasi goroutine leak
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	//go func() {
	//	defer close(destination)
	//	counter := 1
	//	for {
	//		destination <- counter
	//		counter++
	//	}
	//}()

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
	// ada goroutine yang berjalan terus karena terdapat infinite loop dalam fungsi CreateCounter dimana goroutine mengirim data secara terus menerus
}

// context with timeout
// menambahkan sinyal cancel secara otomatis
// cancel akan dieksekusi jika waktunya timeout sudah terlewati
// cocok untuk melakukan query
func CreateCounter2(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) //simulasi slow
			}
		}
	}()
	return destination
}

func TestTimeout(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	// tetap harus memanggil cancel untuk mengatisipasi proses berjalan cepat akan tetapi waktu timeout masih ada
	defer cancel()

	destination := CreateCounter2(ctx)
	fmt.Println("total goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}
	//time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}

// context with deadline
// dapat menentukan waktu deadline misal jam 12 siang
// jika waktu deadline sudah terlewati maka akan mengirimkan sinyal cancel
func TestWithDeadline(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	// tetap harus memanggil cancel untuk mengatisipasi proses berjalan cepat akan tetapi waktu timeout masih ada
	defer cancel()

	destination := CreateCounter2(ctx)
	fmt.Println("total goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}
	//time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}
