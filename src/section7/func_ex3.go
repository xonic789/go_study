package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi")
	prtHello(n - 1)
}

func main() {
	//재귀 함수(Recursion)

	//ex1
	x := fact(7)

	fmt.Println(x)

	//ex2
	prtHello(10)
}
