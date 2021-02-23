package account
import "errors"

type Account struct {
    balance int
    deltaChan chan int
    balanceChan chan int
    errChan chan error
}

func NewAccount(balance int) (a *Account) {
    a = &Account {
        balance: balance,
        deltaChan: make(chan int),
        balanceChan: make(chan int),
        errChan: make(chan error, 1),
    }
    a.run()
    return
}

func (a *Account)run() {
    go func() {
        var amount int
        for {
            select {
                case amount = <-a.deltaChan:
                a.errChan <- a.updateBalance(amount)
                case a.balanceChan <- a.balance:
            }
        }
    }()
}

func (a *Account)updateBalance(amount int) error {
    newBalance := a.balance + amount
    if newBalance < 0 {
        return errors.New("Insufficient funds")
    }
    a.balance = newBalance
    return nil
}

func (a *Account)Balance() int {
    return <-a.balanceChan
}

func (a *Account)Deposit(amount int) error {
    a.deltaChan <- amount
    return <-a.errChan
}

func (a *Account)Withdrawal(amount int) error {
    a.deltaChan <- -amount
    return <-a.errChan
}