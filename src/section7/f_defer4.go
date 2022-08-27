package main

import "fmt"

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b"))
	fmt.Println("in a")
}

func main() {
	//ex1
	a()
}
