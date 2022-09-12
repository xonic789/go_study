package main

import (
	"bufio"
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

	//파일 생성
	file, err := os.Open("sample.csv")
	errCheck1(err)
	//리소스 해제
	defer file.Close()

	//CSV Reader 생성
	//csv.NewReader()
	rr := csv.NewReader(bufio.NewReader(file))

	//CSV 내용 읽기(1)

	row1, err1 := rr.Read()
	errCheck1(err1)
	row2, err2 := rr.Read()
	errCheck1(err2)
	fmt.Println("CSV Read Example")
	fmt.Println(row1[0], row1[1], row1[2], row1[3], row1[1:5])
	fmt.Println(row2[0], row2[1], row2[2], row2[3], row2[1:5])
	//fmt.Println(row)
	fmt.Println("=============================")

	//CSV 내용 읽기(2)
	rows, err := rr.ReadAll()
	errCheck1(err)
	fmt.Println("CSV ReadAll Example")
	//fmt.Println(rows)
	fmt.Println(rows[5][1])
	fmt.Println("=============================")

	// Row 단위로 CSV 파일 읽기

	for i, row := range rows {
		//fmt.Println(i, row)
		for j := range row {
			fmt.Printf("%s      ", rows[i][j])
		}
		fmt.Println()
	}
}
