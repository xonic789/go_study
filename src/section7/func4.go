package main

import "fmt"

func multiply(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 10
	return r1, r2
}

func main() {
	a, b := multiply(10, 5)

	fmt.Println("ex1 : ", a, b)

	// ex2

}
