package main

import "fmt"

func main() {
	//슬라이스(슬라이스 참조 타입)

	//ex1(배열)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println("ex 1 :", arr1)
	fmt.Println("ex 1 :", arr2)
	fmt.Println()

	//ex2(slice)
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex2 : ", slice1)
	fmt.Println("ex2 : ", slice2)

	//ex3(slice exception simulation)
	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	//fmt.Println("ex3 : ", slice3[50]) //길이 index out of range => 길이 만큼만 0으로 초기화

	//ex4(iterate slice)
	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}

}
