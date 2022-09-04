package main

import "fmt"

type test interface{}

func main() {

	//ex1
	/*
		type 인터페이스명 interface {
			메소드1() 반환 값(타입 형)
			메서드2()
		}
	*/

	var t test
	fmt.Println("ex1 : ", t) // empty -> nil

}
