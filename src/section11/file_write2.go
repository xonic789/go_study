package main

import (
	"encoding/csv"
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
	//파일 쓰기
	//CSV 파일 쓰기 예제
	//패키지 저장소를 통해서 Excel

	//파일 생성
	file, err := os.Create("test_write.csv")
	errCheck1(err)

	//리소스 해제
	defer file.Close()

	//csv writer 생성
	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewWriter(file))

	//csv 내용 쓰기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Kee", "4.2"})
	wr.Write([]string{"Lee", "4.1"})
	wr.Write([]string{"Park", "4.7"})
	wr.Write([]string{"Hong", "4.1"})

	wr.Flush() //버퍼 -> 파일로 쓰기

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Println("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name)
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
}
