package dinosaur

import (
  "sync"
)

type Wallet struct {
  balance float64
  mu sync.Mutex
}

func NewWallet() *Wallet {
  return &Wallet{
    balance: InitialBalance,
    mu: sync.Mutex{},
  }
}

func (w *Wallet) Update(value float64) {
  w.mu.Lock()
  defer w.mu.Unlock()

  if value < 0.0 && (w.balance == 0.0 || w.balance + value < 0.0) {
    return
  }
  w.balance += value
}

func (w Wallet) Balance() float64 {
  return w.balance
}
