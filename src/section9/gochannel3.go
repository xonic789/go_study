package main

import (
	"fmt"
	"time"
)

func main() {
	//채널(Channel)
	//ex1(동기 : 버퍼 미사용)

	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			// ch <- true //채널에 데이터 전송(송신) 수신할 때 까지 대기
			ch <- true
			fmt.Println("Go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		// <-ch //채널에서 데이터 수신, 송신할 때 까지 대기
		<-ch
		fmt.Println("Main : ", i)
	}
}
