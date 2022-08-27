package main

import "fmt"

func main() {
	/**
	- 포인터지원 X(파이썬, C#, JAVA 등)
	- 주소 값은 직접 변경 불가능
	- *(애스터리스크)로 사용
	- nil로 초기화 (nil == 0)
	*/

	//ex1
	var a *int
	var b *int = new(int)

	fmt.Println(a)
	fmt.Println(b)

	i := 7

	a = &i // & 주소값 전달
	b = &i

	*a = 77

	fmt.Println("ex1: ", a, &i)
	fmt.Println("ex1: ", &a)
	fmt.Println("ex1: ", *a) //역참조

	fmt.Println()

	fmt.Println("ex1: ", b, &i)
	fmt.Println("ex1: ", &b)
	fmt.Println("ex1: ", *b)
	fmt.Println()
	var c = &i
	d := &i // i의 주소값 전달

	*d = 10 // i의 주소값이 가리키고 있는 메모리 주소에 값 변경

	fmt.Println("ex1: ", c, &i)
	fmt.Println("ex1: ", &c)
	fmt.Println("ex1: ", *c) //역참조

	fmt.Println()

	fmt.Println("ex1: ", d, &i)
	fmt.Println("ex1: ", &d)
	fmt.Println("ex1: ", *d)
}
