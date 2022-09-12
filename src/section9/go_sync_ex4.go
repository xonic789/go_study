package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	/**
	 */

	// 원자성 사용 할 경우 예제
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1) // 고루틴 추가
		go func(n int) {
			//cnt += 1
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1) // 고루틴 추가
		go func(n int) {
			//cnt -= 1
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}
	// Add == Done 횟수 같아야함
	wg.Wait() // 고루틴 종료시까지 대기
	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("WaitGroup End >>>> ", cnt)
	fmt.Println("WaitGroup End >>>> ", finalCnt)
}
