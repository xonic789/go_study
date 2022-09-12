package main

import (
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(%g)은 사용 불가능합니다", f)
	}
	return math.Pow(f, i), nil // 제곱, Nil 반환
}

func main() {
	//ex1
	if f, err := Power(7, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	//ex1
	if f, err := Power(0, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}
}
