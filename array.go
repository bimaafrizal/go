package main

import "fmt"

func main() {
	//daya tampung array tidak bisa bertambah
	var names [3]string = [3]string{"Bima", "Afrizal", "Malna"}
	//bisa juga menggunakan names[0] = "Bima"

	fmt.Println(names[0])
	
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	//function array
	fmt.Println(len(names)) //mencart panjang array
	fmt.Println(len(values))
	//merubah nilai
	values[0] = 100
	fmt.Println(values)
}