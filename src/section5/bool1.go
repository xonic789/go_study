package main

import "fmt"

func main() {
	//암묵적 형 변환 X : 0, Nil -> false 변환 없음.

	//ex 1
	var b1 bool = true
	var b2 bool = false

	b3 := true

	fmt.Println("ex 1 : ", b1)
	fmt.Println("ex 2 : ", b2)
	fmt.Println("ex 3 : ", b3)
}
