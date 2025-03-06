package goroutin

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
// Lock digunakan untuk mengunci data
Lock()
// Unlock digunakan untuk membuka kunci data
Unlock()
// RLock digunakan untuk mengunci data secara bersamaan
RLock()
// RUnlock digunakan untuk membuka kunci data secara bersamaan
RUnlock()
*/
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

/*
hal ini merupakan contoh penggunaan RWMutex
dimana hal ini membuat aman dari perubahan data dan juga pengambilan data secara bersamaan
*/

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}
