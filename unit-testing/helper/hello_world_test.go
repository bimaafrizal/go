package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// untuk menjalankan test ini, ketik di terminal: go test ./...
// untuk melihat tes yang diujikan ketik di terminal: go test -v ./...
// untuk menguji satu function saja, ketik di terminal: go test -v -run TestHelloWorld ./...

// penulisan unit testing menggunakan prefix test
// penulisan function harus diawali dengan Test

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bima")
	if result != "Hello World Bima" {
		//panic("Result must be 'Hello World Bima")
		//jangan gunakan panic karena akan menghentikan proses testing
		//Fail akan menggaggalkan unit test namun akan tetap melanjutkan proses testing
		//FailNow akan menggaggalkan unit test dan menghentikan proses testing
		//Error akan menggaggalkan unit test namun akan tetap melanjutkan proses testing, akan menampilkan errornya (memanggil Fail diakhir)
		//Fatal akan menggaggalkan unit test dan menghentikan proses testing, akan menampilkan errornya (memanggil FatalNow diakhir)
		//t.Fail()
		t.Error("Result must be 'Hello World Bima")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldAfrizal(t *testing.T) {
	result := HelloWorld("Afrizal")
	if result != "Hello World Afrizal" {
		//panic("Result must be 'Hello World Afrizal")
		//t.FailNow()
		t.Fatal("Result must be 'Hello World Afrizal")
	}
	fmt.Println("TestHelloWorldAfrizal Done")
}

//jika menjalankan go test -v -run=TestHelloWorld maka akan menjalankan test yang memiliki prefix TestHelloWorld
//untuk menjalankan semua test yang ada di package helper, ketik di terminal: go test -v ./...

// asertion
// harus import package assertion
// bisa gunakan package testify
// assert jika pengecekan gagal akan memanggil fail
// require jika pengecekan gagal akan memanggil failNow
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bima")
	require.Equal(t, "Hello World Bima", result, "Result must be 'Hello World Bima")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bima")
	assert.Equal(t, "Hello World Bima", result, "Result must be 'Hello World Bima")
	fmt.Println("TestHelloWorldAssert Done")
}

// skip test => membatalkan test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("can't run on windows")
	}

	result := HelloWorld("Bima")
	assert.Equal(t, "Hello World Bima", result, "Result must be 'Hello World Bima")
	fmt.Println("TestHelloWorldAssert Done")
}

// before after
// gunakan testing.M
// wajib membuat function TestMain
// function TestMain akan dijalankan pertama kali sebelum menjalankan test
// hanya berjalan di dalam satu package
func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")

	//runing test
	m.Run()

	//after test
	fmt.Println("After Unit Test")
}

// sub test
// function test di dalam function test
// jika ingin menjalankan salah satu sub test maka ketik di terminal: go test -v -run=TestSubTest/NamaSubTest
// jika ingin menjalankan semua sub test maka ketik di terminal: go test -v -run=TestSubTest
func TestSubTest(t *testing.T) {
	t.Run("Bima", func(t *testing.T) {
		result := HelloWorld("Bima")
		assert.Equal(t, "Hello World Bima", result, "Result must be 'Hello World Bima")
	})
	t.Run("Afrizal", func(t *testing.T) {
		result := HelloWorld("Afrizal")
		assert.Equal(t, "Hello World Afrizal", result, "Result must be 'Hello World Afrizal")
	})
}

// table test
// menyediakan data berupa slice yang berisi parameter dan ekspektasi dari unit test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bima",
			request:  "Bima",
			expected: "Hello World Bima",
		},
		{
			name:     "Afrizal",
			request:  "Afrizal",
			expected: "Hello World Afrizal",
		},
		{
			name:     "Bima Afrizal",
			request:  "Bima Afrizal",
			expected: "Hello World Bima Afrizal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// benchmark
// akan diuji otomatis oleh sistem
// jika ingin menjalankan benchmark dan semua unit test, ketik di terminal: go test -v -bench .
// jika ingin menjalankan benchmark tanpa unit test, ketik di terminal: go test -v -run=NotMathUnitTest -bench .
// menjalankan di root gunakan go test -v -bench=. ./...
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bima")
	}
}

func BenchmarkHelloWorldAfrizal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Afrizal")
	}
}

// sub benchmark
// untuk menjalankan benchmark saja bisa menggunakan go test -v -run=NotMathUnitTest -bench=NameBench/sub
func BenchmarkSub(b *testing.B) {
	b.Run("Bima", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bima")
		}
	})
	b.Run("Afrizal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Afrizal")
		}
	})
}

// table benchmark
func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Bima",
			request: "Bima",
		},
		{
			name:    "Afrizal",
			request: "Afrizal",
		},
	}

	for _, bench := range benchmark {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}
