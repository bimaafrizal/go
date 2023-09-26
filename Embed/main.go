package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

// hasil embed di compile
// hasil embed hasilnya permanent dandata file yang dibaca disimpan dalam binary file golang
// jika sudah dicompile, kita tidak butuh lagi file externalnya dan jika file externalnya dirubah maka isi variabel tidak akan berubah lagi

//go:embed version.txt
var version string

//go:embed uns1.jpeg
var uns1 []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)
	err := ioutil.WriteFile("uns_new.jpeg", uns1, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println("Nama File : ", entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

// hasil yang ditampilkan adalah hasil yang dibaca pada bainary file golang
