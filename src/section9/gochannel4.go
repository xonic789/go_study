package main

import (
	"fmt"
	"runtime"
)

func main() {
	//채널(Channel)
	//ex1(동기 : 버퍼 사용 )
	//버퍼가 가득 차면 송신 블록, 버퍼가 비어 있으면 수신 블록

	runtime.GOMAXPROCS(4)
	ch := make(chan bool, 4) // 버퍼 사용, 버퍼가 가득찰때까지 or 버퍼가 비워질때까지 대기
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			// ch <- true //채널에 데이터 전송(송신) 수신할 때 까지 대기
			ch <- true
			fmt.Println("Go : ", i)
			//time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		// <-ch //채널에서 데이터 수신, 송신할 때 까지 대기
		<-ch
		fmt.Println("Main : ", i)
	}
}
