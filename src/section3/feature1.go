package main

import "fmt"

func main() {
	//GO : 모호하거나, 애매한 문법 금지
	//후치 연산자 허용 i++, 전치 연산자 ++i 비허용
	//증감 연산 반환 값 없음
	//포인터(Pointer 허용, 연산 비허용)

	// ex 1
	sum, i := 0, 0
	for i <= 100 {
		// sum += i++ => X
		sum += i
		i++
	}
	fmt.Println("ex 1 : ", sum)
}
