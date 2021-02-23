package main

import (
    "fmt"
    "Banking/account"
)

func main() {
    acc := account.NewAccount(0)

    fmt.Println("Hello World", acc.Balance())
}