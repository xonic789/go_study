package main

import "fmt"

func main() {
	//짧은 선언
	//함수 안에서만 사용(전역x). 선언 후 할당 예외 발생
	//주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성 높임.

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//shortVar3 := true -> 예외 발생!
	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	//Example
	if i := 10; i < 11 {
		fmt.Println("short Variable Test Success")
	}

}
