package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>", i)
	}
	fmt.Println(name, " end : ", time.Now())
}
func main() {
	//고루틴(Goroutine)
	exe("t1")
	//ex1
	fmt.Println("Main Routine start : ", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine end : ", time.Now())

}
