package goroutin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println("Timer 1 expired", time)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)

	time := <-channel
	fmt.Println("Timer 1 expired", time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timer 1 expired", time.Now())
		group.Done()
	})
	fmt.Println("Waiting for timer", time.Now())
	group.Wait()
}

// Ticker adalah mekanisme yang digunakan untuk mengulang sesuatu secara berkala
// Ticker akan mengirimkan data ke dalam channel setiap interval waktu tertentu
// Ticker akan berhenti ketika dihentikan
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println("Tick", tick)
	}
}
