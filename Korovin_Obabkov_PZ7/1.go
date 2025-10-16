package main

import (
	"fmt"
)

func main() {
	John := BankAccount{"74578235542", "John", 0}
	fmt.Printf("баланс = ", GetBalance(&John))
	fmt.Printf("%v \n", Withdraw(500, &John))
	GetBalance(&John)
	Deposit(&John, 1000)
	Withdraw(500, &John)
}

type BankAccount struct {
	accountNumber, holderName string
	balance                   float64
}

func Deposit(account *BankAccount, amount float64) {
	account.balance += amount
	fmt.Println("Счет пополнен:", account.holderName, GetBalance(account))
}
func GetBalance(account *BankAccount) float64 {
	return account.balance
}
func Withdraw(amount float64, account *BankAccount) error {
	if account.balance >= amount {
		account.balance -= amount
		fmt.Printf("Успешно списано %v. Новый баланс: %v \n", amount, GetBalance(account))
		return nil
	} else {
		return fmt.Errorf("Недостаточно средств!")
	}
}
