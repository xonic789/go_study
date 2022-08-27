package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi!")
	}()
}

func main() {
	//ex1
	sayHello("Golang!")
}
