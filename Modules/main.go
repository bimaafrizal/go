package main

import (
	"fmt"
	say_hello "github.com/bimaafrizal/go-modules/v2"
)

func main() {
	//fmt.Print(Modules.SayHello())
	fmt.Println(say_hello.SayHello("bima"))
	//go get nama module
	//untuk update ketik versi di go.mod kemudian go get
	//jangan lupa update tag agar ikut berubah

	//major update terjadi karena ada perubahan pada isi kode program sehingga tida backward compatible
	//untuk mengatasi kita bisa merubah nama module
	//di go.mod project kita perlu menghapus module yang lama kemudian get kembali
}
