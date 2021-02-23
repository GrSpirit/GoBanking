package account

import (
    "testing"
)

func TestAccountMain(t *testing.T) {
    acc := NewAccount(0)
    acc.Deposit(10000)
    if acc.Balance() != 10000 {
        t.Errorf("Wrong balance. Expected 10000. Got %v", acc.Balance())
    }
}