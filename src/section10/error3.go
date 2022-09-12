package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
	- error 패키지의 New 메소드를 활용한 에러 생성
	- Errorf 보다 더 많이 사용.
	*/

	var err1 error = errors.New("error occurred - 1")
	err2 := errors.New("error occurred - 2")

	fmt.Println("error1:", err1)
	fmt.Println("error1:", err1.Error())

	fmt.Println("error1:", err2)
	fmt.Println("error1:", err2.Error())

}
