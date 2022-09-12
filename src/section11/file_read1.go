package main

import (
	"fmt"
	"os"
)

func errCheck1(err error) {
	if err != nil {
		panic(err)
	}
}

func errCheck2(err error) {
	if err != nil {
		fmt.Println("e")
	}
}

func main() {

	//파일 열기
	file, err := os.Open("sample.txt")
	errCheck1(err)

	//ex1
	fileInfo, err := file.Stat()
	errCheck2(err)
	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) : ", fileInfo, "\n")
	fmt.Println("파일이름(2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일크기(3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일수정시간(4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n\n", ct1)
	//fmt.Println(fd1)
	fmt.Println(string(fd1))

	fmt.Println("=========================================")

	//ex2 (탐색 : Seek(Offset))
	o1, err := file.Seek(20, 0) // 0: 처음위치, 1: 현재위치, 2: 끝위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)

	fmt.Printf("읽기 작업(2) 완료 (%d bytes) (%d ret) \n\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")
	fmt.Println("=========================================")

	//ex3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8)
	errCheck1(err)

	fmt.Printf("읽기 작업(3) 완료 (%d bytes) (%d ret) \n\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")
	fmt.Println("=========================================")

	defer file.Close()
}
