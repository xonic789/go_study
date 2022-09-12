package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Errorf를 이용한 에러 처리

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", a)

	b, err := notZero(0)
	if err != nil {
		log.Fatal(err) // 자동 구현
	}

	fmt.Println("ex2 : ", b)

	fmt.Println("End Error Handling!")
}

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprintf("Hello Golang! %d", n)
		return s, nil
	}
	return "", errors.New(strconv.FormatInt(int64(n), 10) + "은 0입니다")
}
