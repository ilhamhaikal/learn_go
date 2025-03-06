package goroutin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
deadlock adalah kondisi dimana terdapat dua goroutine atau lebih yang saling menunggu
salah satu goroutine menunggu goroutine lain untuk melepaskan lock
sehingga kedua goroutine tersebut tidak akan pernah selesai
dalam contoh ini terdapat dua goroutine yang saling menunggu
yaitu user1 dan user2
*/

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User1", user1.Name)
	user1.Change(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2", user2.Name)
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Ilham",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Yulia",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User1 Balance:", user1.Balance)
	fmt.Println("User2 Balance:", user2.Balance)
}
