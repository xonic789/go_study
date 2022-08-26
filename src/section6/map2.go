package main

import "fmt"

func main() {
	//맵 조회 및 순회(Iterator)

	//ex1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["naver"])
	fmt.Println()

	for k, v := range map1 {
		fmt.Println("ex2 : ", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex2 : ", v)
	}
}
