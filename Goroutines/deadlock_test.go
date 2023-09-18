package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// keadan saling menununggu sehingga tidak ada goroutine yang berjalan

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bima",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Afrizal",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	// saling menunggu sehingga user2 tidak pernah di lock
	// output yang keluar karena fungsi time.Sleep akan tetapi sebenarnya prosesnya belum selesai sehingga transaksi goroutine sebelumnya tidak mengupdate nilai
	// untuk membenarkan fungsi ini, unlock terlebih dahulu user2 pada fungsi balance

	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance", user2.Balance)
}
