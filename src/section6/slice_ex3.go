package main

import "fmt"

func main() {
	//슬라이스 복사
	//copy(복사 대상, 원본)
	//make로 공간을 할당 후 복사 해야 한다.

	//ex1
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)

	//ex2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)

	fmt.Println()

	//ex3

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조 -> 원본 값 변경된다.

	d[1] = 7

	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	fmt.Println()

	//ex4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]

	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)
}
