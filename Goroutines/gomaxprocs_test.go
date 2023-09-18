package Goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// untuk mengetahui jumalh goroutine yang berjalan
// secra default golang menyesuaikan core yang ada
func TestGomaxprox(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoRoutine)

	// merubah banyak thread
	runtime.GOMAXPROCS(20)
	totalThread2 := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread2)
	group.Wait()
}
