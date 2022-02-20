package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int64
	mu      sync.RWMutex
}

func (a *Account) Deposit(value int64) {
	a.mu.Lock()
	a.balance += value
	a.mu.Unlock()
}

func (a *Account) Withdraw(value int64) {
	a.mu.Lock()
	a.balance -= value
	a.mu.Unlock()
}

func (a *Account) Balance() int64 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance
}

// sync.Mutex
// sync.RWMutex
func main() {
	fmt.Println("Go Mutex Example")

	// why is mutex initialization not needed?
	a := &Account{}
	a.Deposit(5)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		a.Withdraw(1)
		time.Sleep(5 * time.Millisecond)
		a.Withdraw(1)
		time.Sleep(5 * time.Millisecond)
		a.Withdraw(1)
	}()

	go func() {
		defer wg.Done()
		a.Deposit(2)
		time.Sleep(5 * time.Millisecond)
		a.Deposit(2)
		time.Sleep(5 * time.Millisecond)
		a.Withdraw(1)
	}()

	wg.Wait()
	fmt.Println(a.Balance())
}
