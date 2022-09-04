package main

import (
	"fmt"
	"time"
)

func work2(v chan int) {
	fmt.Println("Work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ---> ", time.Now())
	v <- 2
}

func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ---> ", time.Now())
	v <- 1
}

func main() {
	/*
		채널(Channel)
		- 고루틴 간에 통신을 위한 메커니즘
		- 고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해서 사용 : 채널(동기식, 래퍼런스 타입)
		- 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
		- 데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고 받을 수 있음)
		- interface{} 전달을 통해서 자료형 상관없이 전송 및 수신가능
		- 값을 전달(복사 후 : bool, int등), 포인터(슬라이스, 맵) 등을 전달시에는 주의 > 동기화 사용(Mutex)
		- 멀티 프로세싱 처리에서 교착상태(경쟁) 주의
		- <-, -> (채널 <- 데이터 : 송신), ( <- 채널 : 수신)
	*/

	//ex1
	fmt.Println("Main : S ---> ", time.Now())
	//var c chan int //채널 선언
	//c = make(chan int)

	v := make(chan int) //채널 선언 및 생성
	go work1(v)
	go work2(v)
	<-v
	<-v
	fmt.Println("Main : E ---> ", time.Now())
}
