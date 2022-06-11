package account

import (
	"math"
	"sync"
)

// Define the Account type here.
type Account struct {
	mu      sync.Mutex // this will ensure ACID of transaction
	cond    bool       //false --> account closed, true --> account open
	balance int64
}

func Open(amount int64) *Account {
	// panic("Please implement the Open function")

	var a Account

	//lock the account before starting the transaction
	a.mu.Lock()
	// unlock the account after transaction
	defer a.mu.Unlock()

	if amount < 0 {
		return nil
	}

	a.balance, a.cond = amount, true

	return &a
}

func (a *Account) Balance() (int64, bool) {
	// panic("Please implement the Balance function")

	//lock the account before starting the transaction
	a.mu.Lock()
	// unlock the account after transaction
	defer a.mu.Unlock()

	return a.balance, a.cond

}

func (a *Account) Deposit(amount int64) (int64, bool) {
	// panic("Please implement the Deposit function")

	//lock the account before starting the transaction
	a.mu.Lock()
	// unlock the account after transaction
	defer a.mu.Unlock()

	//check whether account closed or not
	if !a.cond {
		return a.balance, false
	}

	//not closed, perform transaction
	if amount < 0 && a.balance < int64(math.Abs(float64(amount))) {
		// not enough balance present to perform transaction
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true

}

func (a *Account) Close() (int64, bool) {
	// panic("Please implement the Close function")

	//lock the account before starting the transaction
	a.mu.Lock()
	// unlock the account after transaction
	defer a.mu.Unlock()

	defer a.closeAfterAllTransaction()

	return a.balance, a.cond
}

func (a *Account) closeAfterAllTransaction() {

	a.balance, a.cond = 0, false
}
