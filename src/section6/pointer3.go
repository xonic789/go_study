package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	//포인터 값 전달
	//함수, 메서드 호출 시에 매개변수 값을 복사해서 전달한다. -> 함수, 메서드 내에서는 원본 값 변경 불가능하다.
	//그러므로 원본 값 변경 위해서 포인터 전달.

	//ex1
	var a int = 10
	var b int = 10

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)
	fmt.Println()

	rptc(&a)
	vptc(b)

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)
}
