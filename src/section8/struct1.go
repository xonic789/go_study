package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	//ex1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-902", balance: 12000000}
	park := Account{number: "245-903", interest: 0.025}
	cho := Account{"245-904", 15000000, 0.03}

	fmt.Println("ex1 : ", kim)
	fmt.Printf("ex1 : %p\n", &kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println("ex1 : ", park)
	fmt.Println("ex1 : ", cho)

	fmt.Println()

	//ex2
	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))
}
