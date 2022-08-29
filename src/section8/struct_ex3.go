package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateF(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-901", 20000000, 0.025}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	fmt.Println()

	kim.CalculateD(100000)
	lee.CalculateF(100000) // => 반영됨. 리시버를 연결하면 구조체를 넘기면 알아서 해줌.

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))

}
