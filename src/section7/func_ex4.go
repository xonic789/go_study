package main

import "fmt"

func main() {
	//익명 함수(Anonymous Function)
	//즉시 실행(선언과 동시에)

	//ex1
	func() {
		fmt.Println("ex1 : Anonymous func")
	}()

	//ex2
	msg := "Hello Golang!"

	func(m string) {
		fmt.Println("ex2 : ", m)
	}(msg)

	//ex3
	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20)

	//ex4
	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println("ex4 : ", r)

}
