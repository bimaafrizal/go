package main

import "fmt"

func main() {
	//potongan dari data array
	//ukutan slice bisa berubah
	//slice dan array saling terkait, dimana slice adalah data yang mengakses sebagian atau seluruh data di Array

	//tipe data slice ada 3 yaitu pointer, length, dan capacity
	//pointer adalah posisi awal elemen slice di array
	//length adalah panjang dari slice
	//capacity adalah kapasitas maksimal slice, dimana length tidak boleh melebihi capacity

	//[low,high] => membuat slice dari array dimulai dari index low sampai index high-1
	//[low:] => membuat slice dari array dimulai dari index low sampai index terakhir
	//[:high] => membuat slice dari array dimulai dari index 0 sampai index high-1
	//[:] => membuat slice dari array dimulai dari index 0 sampai index terakhir

	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	//... artinya capacity datanya belum ditentukan

	var slice1 = months[4:7] //pointernya di index 4, lengthnya 3, capacitynya 8 capasitynya 8 karena pounternya di index 4 dan arraynya punya 12 elemen
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//jika merubah data array atau slice maka akan berpengaruh pada data yang lain
	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	//append untuk membuat slice baru untuk menambah data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru
	//make membuat slice baru
	//copy menyalin slice
	slice2 := months[10:]
	fmt.Println(slice2)
	
	slice3 := append(slice2, "Bima")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	//array months tidak berubah karena slice3 sudah membuat array baru akibat dari kapasitas slice sudah penuh

	//membuar slice dari awal
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bima"
	newSlice[1] = "Afrizal"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//hati2 ketika membuat slice, jika [] kosong maka akan menjadi slice jika [5] atau [...] maka akan menjadi array
}