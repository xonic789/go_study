package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

//구조체 Dog 메소드 구현

func (d Dog) bite() {
	fmt.Println(d.name, "Dog bites!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, "Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, "Dog is running")
}

func (d Cat) bite() {
	fmt.Println(d.name, "Cat bites!")
}

func (d Cat) sounds() {
	fmt.Println(d.name, "Cat cries!")
}

func (d Cat) run() {
	fmt.Println(d.name, "Cat is running")
}

type Behavior interface {
	bite()
	sounds()
	run()
}

// 인터페이스를 파라미터로 받아서 실행
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {

	//ex1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	act(dog)
	act(cat)
}
