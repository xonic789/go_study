package main

import "fmt"

type cnt int

func main() {
	//기본 자료형 사용자 정의 타입

	//ex1
	a := cnt(5)
	fmt.Println("ex1 : ", a)

	//ex2
	var b cnt = 15
	fmt.Println("ex1 : ", b)

	//ex3
	// testConverT(b) 예외 발생한다.
	testConverT(int(b)) //형 변환

	testConverD(b)
}

func testConverT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex3 : (Custom Type) : ", i)
}
