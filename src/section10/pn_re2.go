package main

import "fmt"

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!")
}

func main() {

	//ex1
	runFunc()

	fmt.Println("End Main")
}
