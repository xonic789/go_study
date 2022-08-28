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
	//다양한 선언 방법
	//&struct, struct : &struct 포인터를 받아오기 역 참조를 또 하기 때문에 속도는 조금 느리다.
	//인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메서드에 포인터 리시버를 사용할 경우 반드시 &struct

	//선언 방법1
	var kim *Account = new(Account) // 구조체 정보를 new 키워드를 통해 포인터 변수 kim에 담는다.
	kim.number = "245-901"          // kim이 포인터 변수이므로 kim.으로 접근 가능
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방법2
	var hong *Account = &Account{number: "245-902", balance: 150000000, interest: 0.04} // Account를 할당한 메모리 주소를 hong에 할당한다.

	//선언 방법3
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)

	fmt.Printf("ex2 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)

	//ex2
	fmt.Println()

}
