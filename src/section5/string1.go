package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 string = "c:\\go_study\\src\\"
	str2 := `c:\go_study\src\`
	// => str1 == str2

	fmt.Println("ex1 : ", str1)
	fmt.Println("ex1 : ", str2)

	var str3 string = "\ud55c\uae00"

	fmt.Println("ex2 : ", str3)

	fmt.Println("ex3 : ", len(str3))

	fmt.Println("ex4 : ", utf8.RuneCountInString(str3))
	fmt.Println("ex4 : ", len([]rune(str3)))
}
