package main

import "fmt"

func main() {

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15이상")
	}

	if b >= 25 {
		fmt.Println("25이상")
	}

	// syntax 에러 발생 경우
	/*중괄호 생략, 중괄호가 syntax에 맞게 오지 않은 경우
	if a >= 0
	{

	}
	*/
	/*
		if a >= 0
			fmt.Println("Hi")
	*/

	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

}
