package main

import (
	"fmt"
	"section4/lib"
)

func main() {
	var num int32
	fmt.Print("input num : ")
	fmt.Scanf("%d", &num)

	fmt.Println("10보다 큰 수인가 ? : ", lib.CheckNum(num))
}
