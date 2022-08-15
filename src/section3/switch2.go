package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//예제 1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100 미만이다")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50 미만이다")
	default:
		fmt.Println("i -> ", i, " 기본 값")
	}
}
