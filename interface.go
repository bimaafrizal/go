package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

type hitung interface {
	luas() float64
	keliling() float64
}

type Lingkaran struct {
	diameter float64
}

func (l Lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l Lingkaran) keliling() float64 {
	return 3.14 * l.diameter
}

func (l Lingkaran) luas() float64 {
	return 3.14 * l.jariJari() * l.jariJari()
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

type Hewan interface {
	Suara() string
}

func SuaraHewan(h Hewan) {
	fmt.Println("suara hewan itu " + h.Suara())
}

type Kucing struct{}

func (k Kucing) Suara() string {
	return "meong"
}

type Phone interface {
	powerOn() string
	powerOff() string
	volumeUp() string
	volumeDown() string
}
func CekKondisi(hp Phone) string {
	return "cek kondisi hp " + hp.powerOn() + " " + hp.powerOff() + " " + hp.volumeUp() + " " + hp.volumeDown()
}
type Samsung struct {
	user string
}
func (s Samsung) powerOn() string {
	return "power on"
}
func (s Samsung) powerOff() string {
	return "power off"
}
func (s Samsung) volumeUp() string {
	return "volume up"
}
func (s Samsung) volumeDown() string {
	return "volume down"
}

//interface kosong
//interface kosong adalah interface yang tidak memiliki deklarasi method satupun
//apapun data yang kosong
//interface kosong biasanya digunakan sebagai tipe data
//digunakan untuk menampung data dari berbagai tipe data
func Ups(i int) interface{}{
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	//tipe data abstract tidak ada memiliki implementasi langsung
	//berisi definisi method

	//implementasi interface secara otomatis, tidak perlu mengimplementasi interface secara manual
	var bima Person
	bima.Name = "Bima Afrizal"
	SayHello(bima)

	var cat Animal
	cat.Name = "Pus"
	SayHello(cat)

	var bangunDatar hitung
	bangunDatar = persegi{10.0}
	fmt.Println("keliling persegi", bangunDatar.keliling())
	fmt.Println("luas persegi", bangunDatar.luas())

	bangunDatar = Lingkaran{14.0}
	fmt.Println("keliling persegi", bangunDatar.keliling())
	fmt.Println("luas persegi", bangunDatar.luas())

	var kucing Kucing
	SuaraHewan(kucing)

	samsung := Samsung{
		user: "bima",
	}
	fmt.Println(CekKondisi(samsung))


	//untuk menggunakan interface kosong
	var data interface{} = Ups(1)
	fmt.Println(data)
}
