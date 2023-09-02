package main

import "fmt"

func main() {
	// runApplication(10)
	runApp(true)
}

//defer function adalah fungsi yang bisa kita jadwalkan untuk dieksekusi detelah selesai dieksekusi
//akan tetap diekseskusi walaupun terjadi error

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() //akan diekseskusi terakhir
	//harus diawal
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Aplikasi selesai hasilnya", result)
}

//panic function adalah function yang bisa kita gunakan untuk menghentikan program
//panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
//panic function akan menghentikan program

//recover
//recover function adalah function yang bisa kita gunakan untuk menangkap data panic
//recover function harus di deklarasikan dengan defer function karena panic function akan menghentikan program dibawahnya dan menjalankan program diatasnya
func endApp() {
	message := recover() //digunakan untuk menangkap data panic
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
