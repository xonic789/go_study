package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

// bite 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	//ex1
	dog1 := Dog{"poll", 10}
	var inter1 Behavior
	inter1 = dog1
	inter1.bite()

	//ex2
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)

	inter2.bite()

	//ex3
	inters := []Behavior{dog1, dog2}

	//인덱스 형태로 실행
	for idx, _ := range inters {
		inters[idx].bite()
	}

	//값 형태로 실행(인터페이스)
	for _, val := range inters {
		val.bite()
	}

}
