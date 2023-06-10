package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int { return <-balances }
func Withdraw(amount int) bool {
	if Balance() < amount {
		return false
	}
	withdrawals <- amount
	return true
}

func teller() {
	var balance int
	for {
		select {
		case amount := <- deposits:
			balance += amount
		case balances <- balance:
		case amount := <- withdrawals:
			balance -= amount
		}
	}
}

func init() {
	go teller()
}
