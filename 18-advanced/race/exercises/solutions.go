// Package exercises 提供 Data Race 章节代码练习的参考实现。
package exercises

import "sync"

// Account 使用同一临界区完成余额检查和扣减。
type Account struct {
	mu      sync.Mutex
	balance int
}

// NewAccount 创建指定初始余额的账户。
func NewAccount(balance int) *Account {
	return &Account{balance: balance}
}

// Withdraw 在余额充足时扣减金额。
func (a *Account) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 || a.balance < amount {
		return false
	}
	a.balance -= amount
	return true
}

// Balance 返回当前余额。
func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// SignalAfterWrite 写入共享值后发送通知，channel 建立写入到读取的 happens-before 关系。
func SignalAfterWrite(target *int, value int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		*target = value
		close(done)
	}()
	return done
}
