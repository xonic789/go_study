package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	/*
		고루틴 동기화 객체
		- 동기화 상태(조건) 메소드 사용
		- Wait, notify, notifyAll : 기타 언어
		- Wait, Signal, Broadcast : Go 언어
		-  => 실행 흐름 제어 가능
	*/

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan int, 5) //비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting: ", n)
			cond.Wait() //실행 흐름 대기
			fmt.Println("Goroutine Wake End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		//<-c
		fmt.Println("received : ", <-c)
	}
	/*

		for i := 0; i < 5; i++ {
			mutex.Lock()
			fmt.Println("Wake Goroutine(Signal) : ", i)
			cond.Signal() //한 개씩 깨움(모든 고루틴 생성 후)
			mutex.Unlock()
		}
	*/

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	cond.Broadcast() //모두 깨움
	mutex.Unlock()

	time.Sleep(2 * time.Second)

}
