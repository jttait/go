package main

import (
	"fmt"
	
	"github.com/jttait/gopl.io/ch9/bank"
)

func main() {
	fmt.Println(bank.Balance())
	bank.Deposit(100)
	fmt.Println(bank.Balance())
	bank.Withdraw(100)
	fmt.Println(bank.Balance())
	bank.Withdraw(100)
	fmt.Println(bank.Balance())
}
