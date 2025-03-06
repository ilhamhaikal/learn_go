package goroutin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
pool adalah package yang digunakan untuk mengelola data yang sudah tidak terpakai
*/

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Ilham")
	pool.Put("Muhammad")
	pool.Put("Haikal")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Done")
}
