package main

import "fmt"

func main() {
	var num1 float32 = 0.14
	var num2 float32 = .75644
	var num3 float32 = 442.012314
	var num4 float32 = 10.0

	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex1 : ", num2)
	fmt.Println("ex1 : ", num3)
	fmt.Println("ex1 : ", num4)
	fmt.Println("ex1 : ", num4-0.1)
	fmt.Println("ex1 : ", float32(num4-0.1))
	fmt.Println("ex1 : ", float64(num4-0.1))
	fmt.Println("ex1 : ", num5)
	fmt.Println("ex1 : ", num6)
	fmt.Println("ex1 : ", num7)

}
