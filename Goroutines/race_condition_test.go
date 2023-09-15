package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// race condition adalah kondisi dimana 2 goroutine atau lebih mengakses data yang sama pada waktu yang bersamaan
// akan sangat bahaya ketika manipulasi data yang sama

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	// hasilnya akan berbeda karena ada titik dimana goroutine mengambil nilai yang sama
	time.Sleep(5 * time.Second)
}

// mutex
// mutex bisa digunakan untuk locking dan unlocking data
// hanya 1 goroutine yang bisa lock data yang sama
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t.Log("Counter = ", x)
}

// RW Mutex
// jika mutex saja akan rebutan antara proses membaca dan merubah data
// digunakan untuk lock read dan lock untuk write
// khusus untuk struct yang diakses oleh beberapa goroutine
type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.balance = account.balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total balance", account.GetBalance())
}
