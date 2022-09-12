package main

import (
	"fmt"
	oper "section12/arithmetic"
)

func main() {
	nums := oper.Numbers{X: 100, Y: 10}
	fmt.Println("Package Used(1) : ", nums.Plus())
	fmt.Println("Package Used(2) : ", nums.Minus())
	fmt.Println("Package Used(3) : ", nums.Multi())
	fmt.Println("Package Used(4) : ", nums.Divide())
	fmt.Println("Package Used(5) : ", nums.SquarePlus())
	fmt.Println("Package Used(6) : ", nums.SquareMinus())
}
