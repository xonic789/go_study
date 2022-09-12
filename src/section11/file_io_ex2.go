package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	//bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) // Writer 반환(버퍼사용)
	wt.WriteString("Hello Golang!\n File Write Test1\n")
	wt.Write([]byte("Hello Golang!\n File Write Test2\n"))

	errCheck(err)

	//버퍼 정보 출력
	fmt.Printf("사용한 Buffer Size (%d bytes)\n\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d bytes)\n\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d bytes)\n\n", wt.Size())

	wt.Flush()
	fmt.Println("쓰기 작업 완료")
	fmt.Println("===============================")

	rt := bufio.NewReader(file)
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)
	data, _ := rt.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("===============================")
	fmt.Println(string(b))
	fmt.Println("===============================")

	file.Close()

}
