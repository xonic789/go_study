package main

import "fmt"

func main() {
	//ex 1

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	//에러 발생1
	// for i := 0; i < 5; i++
	// {

	// }

	//ex 2
	// for {
	// 	fmt.Println("ex2 : Hello, Golang!")
	// 	fmt.Println("ex2 : Infinite Loop!")
	// }

	//ex 3
	loc := []string{"Seoul", "Busan", "Inchoen"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
