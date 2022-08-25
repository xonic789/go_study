package main

import (
	"fmt"
	"strings"
)

func main() {
	//ex1(결합 : 일반 연산)
	str1 := "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from" +
		"programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are" +
		"written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other" +
		"words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go,"

	str2 := "such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand."

	fmt.Println("ex1 : ", str1+str2)
	//ex2(결합 : Join)
	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, ""))
}
