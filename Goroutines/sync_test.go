package Goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// wait group digunakan untuk menunggu proses goroutine selesai dilakukan, sebelum aplikasi selesai
// jangan gunakan sleep
// untuk menandai ada proses goroutine kita bisa menggunakan method add, setelah selesai kita bisa gunakan method done
// untuk menunggu kita bisa menggunakan method wait

func RunAsync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

// once
// fungsi yang hanya bisa dipakai sekali
// memastikan goroutine pertama yang bisa mengeksekusi functionnya
var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
}

// pool
// digunakan untuk menyimpan data, kita bisa mengambil data dari pool kemudian kita kembalikan lagi
// untuk manage data database
// aman dari race condition

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// make new default value
		New: func() interface{} {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Bima")
	pool.Put("Afrizal")
	pool.Put("Malna")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			//mengambil data pool
			data := pool.Get()
			fmt.Println(data)
			//mengembalikan data pool
			pool.Put(data)
		}()
		group.Done()
	}

	group.Wait()
	//time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}

// map
// map aman untuk menggunakan concurrent
// store untuk menyimpan
// load untuk mengambil data
// delete untuk mengahapus data
// range untuk melakukan iterasi di seluruh data di map
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

// cond
// locking berbasis kondisi, untuk melihat apakah perlu mengunggu atau tidak
// membutuhkan locker mutex
// terdapat fungsi Wait untuk menunggu
// fungsi Signal digunakan agar goroutine tidak perlu menunggu lagi
// fungsi broadcast digunbakan agar semua goroutine tidak perlu menunggu lagi
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	//menanyakan apakah perlu wait atau tidak
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cond.Broadcast()
	//	}
	//}()

	group.Wait()
}

// atomic
// menggunakan data primitive secara aman pada concurrent
// tidak perlu locking
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {

				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("hasilnya", x)
}
