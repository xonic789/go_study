package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//기본적인 메서드 에러 처리 예제
	f, err := os.Open("unnamedfile") //err
	if err != nil {
		//log.Fatal(err.Error())
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", f.Name())
}
