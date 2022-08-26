package main

import "fmt"

func main() {
	//맵 조회 및 순회(Iterator)

	//ex1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}

	fmt.Println("ex1 : ", map1)
	fmt.Println()

	map1["home2"] = "http://test2.com"
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	map1["home2"] = "http://test2-2.com"

	delete(map1, "home2")
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	delete(map1, "google")
	fmt.Println("ex1 : ", map1)
}
