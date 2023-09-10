package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Looping For")

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// for dengan statement
	//init statement statement sebelum perulangan
	//post statement statement setelah perulangan
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Bima", "Afrzal", "Rahman"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//foreach/forange
	for i, name := range slice {
		fmt.Println("Index", i, "=", name)
	}
	//jika tidak ingin menggunakan index pakai _
	// for _, name := range slice {
	// 	fmt.Println(name)
	// }

	person := make(map[string]string)
	person["name"] = "Bima"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
	//jika map berupa key and value

	//break and continue
	//break untuk menghentikan seluruh perulangan
	//continue untuk menghentikan perulangan saat ini, dan melanjutkan ke perulangan berikutnya
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
			// continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
