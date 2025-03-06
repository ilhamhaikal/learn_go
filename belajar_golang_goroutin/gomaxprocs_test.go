package goroutin_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Run asyncronous")
			group.Done()
		}()
	}

	TotalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", TotalCPU)

	// runtime.GOMAXPROCS(20)
	TotalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", TotalThread)

	TotalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", TotalGoroutine)

}
