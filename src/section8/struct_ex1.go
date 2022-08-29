package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest}
}

func main() {
	//구조체 생성자 패턴 예제

	//ex1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "245-902"
	lee.balance = 130000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	//ex2
	park := NewAccount("245-903", 170000000, 0.04)
	fmt.Println("ex2 : ", park)
}
