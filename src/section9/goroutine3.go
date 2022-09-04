package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>>", r, i)
	}
	fmt.Println(name, " end : ", time.Now())

}

func main() {
	//고루틴(Goroutine)
	//멀티 코어(다중 cpu) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                        // 현재 컴퓨터의 cpu 코어 수를 가져와서 사용
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) //설정 값 출력

	//ex1
	fmt.Println("Main Routine start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i) //고루틴 100개 생성
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine end : ", time.Now())

}
