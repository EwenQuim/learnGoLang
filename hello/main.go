package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
	slices()
	structs()
}

func slices() {
	xs := []int{}
	xs = append(xs, 1)
	xs = append(xs, 2)
	xs = append(xs, 3)
	for i, x := range xs {
		fmt.Printf("xs[%d] = %d\n", i, x)
	}
}

func structs() {
	a := Account{id: 3}
	fmt.Println(a.String())
}

type Account struct {
	id      int
	balance float64
}

func (a *Account) String() string {
	return fmt.Sprintf("Account[%d]: $%0.2f", a.id, a.Balance())
}

func (a *Account) Deposit(amt float64) float64 {
	a.balance += amt
	return a.balance
}

func (a *Account) Withdraw(amt float64) float64 {
	a.balance -= amt
	return a.balance
}

func (a *Account) Balance() float64 {
	return a.balance
}
