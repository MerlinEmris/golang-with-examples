package main

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	Mutex   *sync.Mutex
}

func (a *Account) WithDraw(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance -= value
	a.Mutex.Unlock()
	wg.Done()
}

func (a *Account) Deposit(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance += value
	a.Mutex.Unlock()
	wg.Done()
}

func main() {
	var m sync.Mutex

	account := Account{
		Balance: 1001,
		Mutex:   &m,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go account.WithDraw(300, &wg)
	go account.Deposit(100, &wg)
	wg.Wait()

	fmt.Println("Balance updated")
	fmt.Println(account.Balance)
}
