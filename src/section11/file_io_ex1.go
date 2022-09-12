package main

import (
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용

	s := "Hello Golang!\n File Write Test\n"

	//파일 모드(chmod, chown, chgrp) -> 퍼미션
	//읽기 : 4, 쓰기 : 2, 실행 : 1
	//소유자, 그룹, 기타 사용자 순서(644)
	//https://golang.org/pkg/os/#FileMode
	err := os.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := os.ReadFile("sample.txt")
	errCheck(err)

	fmt.Println("===============================")
	fmt.Println(string(data))
	fmt.Println("===============================")

}
