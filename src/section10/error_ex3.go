package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// PowError 에러처리 구조체
type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 파라미터
	message string    // 에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time.Now(), f, "0은 사용 불가능합니다"}
	}
	if math.IsNaN(f) {
		return 0, PowError{time.Now(), f, "숫자가 아닙니다"}
	}
	if math.IsNaN(i) {
		return 0, PowError{time.Now(), i, "숫자가 아닙니다"}
	}
	return math.Pow(f, i), nil
}

func main() {

	//ex1
	v, err := Power(10, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", v)

	//ex2
	t, err := Power(0, 3)
	if err != nil {
		//log.Fatal(err.Error())
		fmt.Println(err.(PowError).time)
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
	}

	fmt.Println("ex2 : ", t)
}
