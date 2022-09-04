package main

import "fmt"

func main() {
	//실제 타입 검사 switch
	//빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능

	//ex1
	checkType(true)
	checkType(1)
	checkType(22.555)
	checkType(nil)
	checkType("Hello Golang")
}

//func checkType(arg interface{}) {
//	switch arg.(type) {
//	case Account:
//	case *Account:
//	default:
//	}
//}

func checkType(arg interface{}) {
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool type", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int type", arg)
	case float64:
		fmt.Println("This is a float64 type", arg)
	case string:
		fmt.Println("This is a string type", arg)
	case nil:
		fmt.Println("This is a nil type", arg)
	default:
		fmt.Println("Unknown type", arg)
	}
}
