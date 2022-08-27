package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a int, b int) {
	fmt.Println("ex1 : ", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_references(i *int) {
	*i *= 10
}

func main() {
	//함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)

	//ex1
	sum(100, add)

	//ex2
	//call by value
	a := 100
	multi_value(a)
	fmt.Println("ex2 : ", a)

	//ex3
	b := 100
	multi_references(&b)
	fmt.Println("ex3 : ", b)
}
