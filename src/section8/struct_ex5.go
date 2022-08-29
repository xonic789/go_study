package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {
	/**
	- 구조체 임베디드 메소드 오버라이딩 패턴
	*/

	//ex1
	ep1 := Employee{"kim", 20000000, 150000}
	ep2 := Employee{"park", 15000000, 200000}

	ex := Executives{
		Employee{"lee", 50000000, 1000000},
		100000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	fmt.Println("ex2 : ", int(ex.Calculate()))
	fmt.Println("ex2 : ", ex.Employee.Calculate()+ex.specialBonus)

}
