package main

import "fmt"

func main() {
	//ex 1
	a := 30 / 15
	switch a {
	case 2, 4, 6: //i가 2, 4, 6일 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: //i가 1, 3, 5인 경우
		fmt.Println("a -> ", a, "는 홀수")
	}

	//ex 2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c":
		fmt.Println("c")
	}
}
