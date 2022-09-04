package main

import "fmt"

func main() {
	//인터페이스 활용(빈 인터페이스)
	//빈 인터페이스 : 함수 매개변수, 리턴 값, 구조체 필드 등으로 사용가능 -> 어떤 타입으로 변환 가능
	//모든 타입을 나타내기 위해 빈 인터페이스 사용
	//동적 타입으로 생각하면 쉽다 (타 언어의 Object 타입)

	//ex1
	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "Golang!"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)
}

func printContents(v interface{}) {
	fmt.Printf("Type : (%T)", v) //원본 타입
	fmt.Println("ex1 : ", v)
}
