package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test_write_txt")
	errCheck1(err)

	//리소스 해제
	defer file.Close()

	//쓰기 예제 1
	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write([]byte(s1))
	errCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes)\n", n1)

	file.Sync() //Write Commit(Stable)

	//쓰기 예제 2
	s2 := "\nHello Golang!\n File Write Test ! -1 \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Printf("쓰기 작업(2) 완료 (%d bytes)\n", n2)

	file.Sync() //Write Commit(Stable)

	//쓰기 예제 3
	s3 := "\nTest WriteAt ! - 2\n"
	n3, err := file.WriteAt([]byte(s3), 100) //Len(offset) 조절하면서 테스트

	fmt.Printf("쓰기 작업(3) 완료 (%d bytes)\n", n3)

	file.Sync() //Write Commit(Stable)

	n4, err := file.WriteString("Hello Golang! \n File Write Test! - 3 \n")
	errCheck2(err)
	fmt.Printf("쓰기 작업(4) 완료 (%d bytes)\n", n4)
}

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
