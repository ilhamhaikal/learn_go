package goroutin

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Hello, Ilham"
		fmt.Println("Data has been sent to channel")
	}()

	data := <-channel
	fmt.Println("Data received from channel:", data)
	time.Sleep(5 * time.Second)
}

func GiveMeSomeData(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello, Ilham"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeSomeData(channel)

	data := <-channel
	fmt.Println("Data received from channel:", data)
	time.Sleep(5 * time.Second)
}

// channel in function adalah parameter yang digunakan untuk mengirim data dari goroutine yang satu ke goroutine yang lain
func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello, Ilham"
}

// channel out function adalah parameter yang digunakan untuk menerima data dari goroutine yang satu ke goroutine yang lain
// hanya bisa menerima data dari channel, jika memaksa mengirim data ke channel maka akan terjadi error
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Data received from channel:", data)
}

// test channel in and out
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	data := <-channel
	fmt.Println("Data received from channel:", data)
	time.Sleep(5 * time.Second)
}

// buffered channel adalah channel yang memiliki kapasitas lebih dari 0
// buffered channel memungkinkan kita untuk mengirim data sebanyak kapasitas channel tanpa harus menunggu data tersebut dibaca
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Hello, Ilham"
		channel <- "Hello, Muhammad"
		channel <- "Hello, Haikal"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

// range channel digunakan untuk mengambil data dari channel secara berulang
// range channel akan berhenti ketika channel sudah ditutup
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data:", data)
	}

	fmt.Println("Selesai, See you next time")
}

// select adalah statement yang digunakan untuk mengambil data dari beberapa channel secara bersamaan
// select akan memilih channel yang siap mengirimkan data terlebih dahulu
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeSomeData(channel1)
	go GiveMeSomeData(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++

		// default case akan dijalankan jika tidak ada channel yang siap
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
