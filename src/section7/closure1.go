package main

import "fmt"

func main() {
	/**
	- 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	- 함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언 된 변수에 접근 가능한 함수
	- 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	- 함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트, 무분별한 전역변수 남발...

	*/

	//ex1
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println("ex1 : ", r1)

	m, n := 5, 10            // 변수가 캡쳐 (선언과 동시에 캡쳐됨.)
	sum := func(c int) int { //익명함수 선언과 동시에 할당
		return m + n + c // 지역변수 소멸 되지 않음. (함수 호출시마다 사용 가능)
	}

	r2 := sum(10)

	fmt.Println("ex2 : ", r2)
}
