package main

import "fmt"

type Accountable interface {
	Pay(amountToPay float32)
	SetNext(account Accountable)
	CanPay(amount float32) bool
}

type Account struct {
	nextAcc Accountable
	balance float32
}

func (a *Account) SetNext(account Accountable) {
	a.nextAcc = account
}

func (a *Account) CanPay(amount float32) bool {
	return amount <= a.balance
}

func (a *Account) Pay(amountToPay float32) {
	if a.CanPay(amountToPay) {
		fmt.Println("Payment complete")
	} else if a.nextAcc != nil {
		fmt.Println("Not enough balance, trying next account...")
		a.nextAcc.Pay(amountToPay)
	} else {
		fmt.Println("Not enough balance on all accounts!")
	}
}

type Bank struct {
	Account
}

type Qiwi struct {
	Account
}

type Bitcoin struct {
	Account
}

func main() {
	bank := Bank{Account{balance: 100}}
	qiwi := Qiwi{Account{balance: 200}}
	bitcoin := Bitcoin{Account{balance: 300}}

	bank.SetNext(&qiwi)
	qiwi.SetNext(&bitcoin)

	bank.Pay(250)
}
