package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num++
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 예제
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요
	//뮤텍스 (Mutex) : 여러 고루틴에서 작업하는 공유 데이터 보호
	//sync.Mutex 선언 후 Lock, Unlock 사용

	//동기화 사용하지 않은 경우 예제
	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func(n int) {
			c.increment()
			done <- true
			runtime.Gosched() // cpu 양보
		}(i)
	}

	for i := 1; i <= 10000; i++ {
		<-done

	}
	c.result()
}
