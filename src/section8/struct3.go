package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {
	//ex1
	c1 := car{"red", "22d"} //사용자 정의 타입 변수 할당
	c2 := new(car)          // 포인터형 변수에 구조체 할당
	c2.color, c2.name = "white", "sonata"
	c3 := &car{}                // 메모리 주소를 할당. (c3에 할당되는 건 메모리 주소)
	c4 := &car{"balck", "520d"} // 메모리 주소를 할당

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
}
