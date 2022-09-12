package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func onceTest() {
	fmt.Println("Once Test Execute")
}

func main() {
	/**
	### 고루틴 동기화 고급
		- Once : 한 번만 실행(주로 초기화에 사용)
		- Do로 실행
	*/

	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once) // once 객체 생성

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Goroutine : ", n)
			once.Do(onceTest)
		}(i)

	}

	time.Sleep(2 * time.Second)
}
