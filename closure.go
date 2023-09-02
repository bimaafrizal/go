package main

import "fmt"

func main() {
	//kemampuan sebuah fungsi berinteraksi dengan data disekitar dengan scope yang sama
	counter := 0
	// name := "Bima"
	
	incerement := func() {
		// name = "Bima Adi"
		//jika begini maka nama akan berubah menjadi Bima Adi
		//namun jika tidak maka akan tetap Bima
		//name := "Afrizal"
		fmt.Println("increment")
		counter++
	}

	incerement();
	incerement();
	fmt.Println(counter)
	
}