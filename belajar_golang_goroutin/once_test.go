package goroutin

import (
	"fmt"
	"sync"
	"testing"
)

/*
once adalah package yang digunakan untuk memastikan bahwa suatu fungsi hanya dijalankan sekali
meskipun fungsi tersebut dipanggil berkali-kali
*/

var counter = 0

func OnlyOunce() {
	counter++
}

func TestOnlyOunce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOunce) // hanya memasukan fungsi yang tidak memiliki parameter
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter:", counter)
}
