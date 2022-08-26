package main

import "fmt"

func main() {

	//ex1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok) //두 번째 리턴 값으로 키 존재 유무 확인

	//ex2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exists")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exists")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exists")
	}
}
