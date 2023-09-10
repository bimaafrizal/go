package main

import "fmt"

func main() {
	sayHello("bima", "afrzal")
	firstName := "Bima"
	sayHello(firstName, "afrzal")
	fmt.Println(getHello("Bima"))
	fmt.Println(getHello(""))

	fmt.Println(getFullName("Bima", "Afrzal"))
	//menyimpan data return value
	firstName2, lastName2 := getFullName("Bima2", "Afrzal2")
	fmt.Println(firstName2, lastName2)

	//menghiraukan return value
	firstName3, _ := getFullName2("Bima3", "Afrzal3")
	//gunakan _ jika tidak ingin memanggil data return value
	fmt.Println(firstName3)

	a, b, c := getFullName3()
	fmt.Println("named return value:", a, b, c)

	//variadic
	fmt.Println(sumAll(1, 4, 5, 5, 6))
	//variadic dengan slice parameter
	slice := []int{10, 20, 40, 50}
	total := sumAll(slice...)
	fmt.Println(total)

	//function value
	sayGoodBye := getGoodBye
	result := sayGoodBye("Bima")
	fmt.Println(result)
	//variabel sayGoodBye memiliki value function getGoodBye

	//function value
	sayHelloWithFilter("bima", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//anonymous function
	blacklist := func(name string) bool {
		return name == "admin"
	}
	fmt.Println("Anonymous Function")
	registerUser("admin", blacklist)
	registerUser("bima", blacklist)

	//recrusive
	fmt.Println(faktorial(3))
	fmt.Println(factorialRecrusive(5))
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello World", firstName, lastName)
}

// function return value
func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	}
	return "Hello " + name
}

// function return multiple value
func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

//menghiraukan return value
func getFullName2(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

//named return values
//membuat variabel secara langsung di tipe data return functionnya
func getFullName3() (firstName, middleName, lastName string) {
	firstName = "Bima"
	middleName = "Afrizal"
	lastName = "Malna"
	return
	//tidak perlu menulis parameter pada return karena sudah deklarasi variabel di atas
}

//variadic function
//parameter yang berada di posisi terakhir memiliki kemampuan dijadikan sebuah varargs
//varargs => bisa menerima lebih dari satu input atau bisa disebut array
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//function value
//bisa dibuat sebagai tipe data, bisa disimpan di dalam variabel, bisa dibuat indepent tanpa group
func getGoodBye(name string) string {
	return "Good Bye " + name
}

//function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	namedFiltered := filter(name)
	fmt.Println("Hello", namedFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	}
	return name
}

//func type declaration
//agar tidak terlalu panjang ketika menulis parameter
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	namedFiltered := filter(name)
	fmt.Println("Hello", namedFiltered)
}

//anonymous function
//membuat function secara langsung di variabel atau parameter tanpa harus membuat function terlebih dahulu
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	}
	fmt.Println("Welcome", name)
}

//recrusive function
//function yang memanggil dirinya sendiri
func faktorial(nilai int) int {
	result := 1
	for i := nilai; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}
