package main

import "fmt"

func main() {
	//채널(Channel)
	//Close : 채널 닫기(수신자 측에서 확인) 주의 -> 닫힌 채널에 값 전송시 패닉(에러) 발생
	//Range : 채널 안에서 값을 꺼낸다(순회), 채널 닫아야(Close) 반복문 종료 -> 채널이 열려 있고 값 전송하지 않으면 계속 대기

	//ex1
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
		//수신하지 않았으므로 close 쓰지 않으면 계속 대기한다.
	}()

	for i := range ch { // 채널에서 값을 꺼내온다.(채널이 Close 될 때까지)
		fmt.Println("ex1 : ", i)
	}
}
