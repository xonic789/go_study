// 사용자 패키지 설치 및 활용 예제
package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 외부 저장소 패키지 설치
	// 2가지 방법
	//1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	//2. go get 패키지 주소 설치 -> 선언

	xfile := "sample.xlsx"

	file, err := excelize.OpenFile(xfile)

	if err != nil {
		panic("Excel Loads Error")
	}

	for _, sheetName := range file.GetSheetList() {
		rows, err := file.GetRows(sheetName)
		if err != nil {
			panic("Excel Rows Error")
		}
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Printf("%s\t", colCell)
			}
			fmt.Println()
		}
	}
}
