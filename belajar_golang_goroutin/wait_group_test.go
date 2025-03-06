package goroutin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
wait group adalah package yang digunakan untuk menunggu goroutine selesai
wait group digunakan ketika kita ingin menunggu beberapa goroutine selesai
*/

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	fmt.Println("Run asyncronous")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(group)
	}

	group.Wait()
	fmt.Println("Complete")
}
