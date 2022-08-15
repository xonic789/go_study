package main

import "fmt"

func main() {
	//ex 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("ex1 : ", sum1)

	//ex 2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
		////후치연산 반환값 x
	}

	fmt.Println("ex2 : ", sum2)

	//ex 3
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++

	}
	fmt.Println("ex3 : ", sum3)

	//ex 4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex 4 : ", i, j)
	}

}
