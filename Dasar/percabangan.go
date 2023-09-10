package main

import "fmt"

func main() {
	name := "Afrzal"

	if name == "Bima2" {
		fmt.Println("Hello Bima")
	} else if name == "Afrzal" {
		fmt.Println("Hello Afrzal")
	} else {
		fmt.Println("Hello Stranger")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

	//switch
	switch name {
	case "Bima":
		fmt.Println("Hello Bima")
	case "Afrzal":
		fmt.Println("Hello Afrzal")
	default:
		fmt.Println("Hello Stranger")
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	}
}
