package main

import "fmt"

func main() {
	//ex1
	i := 7
	p := &i

	fmt.Println("ex1 : ", i, *p, &i, p)

	*p++

	fmt.Println("ex1 : ", i, *p, &i, p)

	*p = 7777

	fmt.Println("ex1 : ", i, *p, &i, p)

	i = 77

	fmt.Println("ex1 : ", i, *p, &i, p)
}
