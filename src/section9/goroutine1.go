package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	//고루틴(Goroutine)
	//코어 수에 상관없이 동시에 실행되는 것처럼 보이는 것
	//실제로는 코어 수에 비례해서 동시에 실행됨
	//비동기적으로 동작하는 함수 루틴
	//채널을 통한 통신으로 동기화 가능
	//멀티 코어 CPU에서 각 코어는 별도의 고루틴을 동시에 실행
	//고루틴은 메인 함수가 종료되어도 계속 실행됨
	//주의 : 고루틴은 메인 함수가 종료되면 같이 종료됨
	//주의 : 고루틴은 동시에 실행되는 것처럼 보이지만 실제로는 순차적으로 실행됨
	//주의 : 고루틴은 메모리 공유 시에 문제 발생 -> 고루틴 동기화 필요
	//주의 : 고루틴은 많은 메모리를 사용 -> 고루틴 사용 시 메모리 사용량 고려

	//멀티 쓰레드 장점과 단점
	//장점 : 응답성이 좋아짐, 자원 공유가 용이, 작업이 분리되어 코드 간결
	//단점 : 구현하기 어려움(복잡성이 증가), 테스트 및 디버깅 어려움, 동기화 처리 필요, 자원 공유 시 문제 발생 가능성 증가
	//고루틴 장점과 단점
	//장점 : 매우 가볍고 간단, 자원 공유가 용이
	//단점 : 응답성이 떨어짐, 자원 공유 시 문제 발생 가능성 증가

	exe1() // 가장 먼저 실행(일반적인 실행 흐름)

	fmt.Println("Main Routine Start", time.Now())
	go exe2() // 고루틴 실행(비동기적으로 실행)
	go exe3()
	time.Sleep(4 * time.Second) // 4초 대기
	fmt.Println("Main Routine End", time.Now())
}
