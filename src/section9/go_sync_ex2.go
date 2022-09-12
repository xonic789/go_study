package main

import (
	"fmt"
	"sync"
)

func main() {
	/**
	### 고루틴 동기화 고급
	- 대기 그룹
	- 실행 흐름 변경(고루틴이 종료 될 때까지 대기 가능)
	- WaitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시까지 대기)
	- Add로 추가 된 고루틴 개수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다.
	*/

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1) // 고루틴 추가
		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}
	// Add == Done 횟수 같아야함
	wg.Wait() // 고루틴 종료시까지 대기
	fmt.Println("WaitGroup End")
}
