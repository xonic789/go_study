package main

import "fmt"

func main() {
	//채널(Channel)
	//채널 또한 함수의 반환 값으로 사용 가능

	//ex1
	c := sum(100)

	fmt.Println("ex1 : ", <-c)
}

func sum(cnt int) <-chan int {
	sum := 0
	out := make(chan int)

	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		out <- sum
	}()

	return out
}
