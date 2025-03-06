package goroutin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Mutex adalah package yang digunakan untuk menghandel race condition
dimana pada race condition terjadi ketika ada dua goroutine atau lebih yang mengakses data yang sama secara bersamaan
sehingga terjadi konflik pada data tersebut
solusi dari situasi tersebut adalah menggunakan mutex
mutex adalah lock yang digunakan untuk mengunci data yang sedang diakses oleh goroutine
sehingga goroutine lain tidak bisa mengakses data tersebut
*/
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
}
