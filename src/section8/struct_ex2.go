package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	//ex1
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-901", 20000000, 0.025}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))

}
