package belajargolangcontex

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
context adalah package yang digunakan untuk mengatur proses concurency
immutable pada context artinya tidak bisa diubah
context bisa dibuat dengan menggunakan context.Background()
context bisa dibuat dengan menggunakan context.TODO()
*/

func TestContex(t *testing.T) {
	background := context.Background()
	fmt.Println("Background", background)

	todo := context.TODO()
	fmt.Println("TODO", todo)
}

/*
context with value adalah context yang bisa menyimpan data
data yang disimpan harus berupa key value

goroutin leak adalah goroutine yang tidak selesai dan tidak dihentikan
*/
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b")) // tidak bisa mengakses value dari parent maka nilanya akan nil
	fmt.Println(contextF.Value("d"))
	fmt.Println(contextF.Value("e"))

}

func CreatCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextCancel(t *testing.T) {
	fmt.Println("mulai goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreatCounter(ctx)

	fmt.Println("Setelah goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("goroutine selesai", runtime.NumGoroutine())
}

func TestWithTimeOut(t *testing.T) {
	fmt.Println("mulai goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreatCounter(ctx)

	fmt.Println("Setelah goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println(n)
	}
	// mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("goroutine selesai", runtime.NumGoroutine())
}

func TestWithDeadline(t *testing.T) {
	fmt.Println("mulai goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreatCounter(ctx)

	fmt.Println("Setelah goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println(n)
	}
	// mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("goroutine selesai", runtime.NumGoroutine())
}
