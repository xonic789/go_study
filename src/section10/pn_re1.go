package main

import "fmt"

func main() {

	fmt.Println("Start Main")
	panic("Error occurred : user stopped!") // 방법1
	//log.Panic("Error occurred : user stopped!") // 방법2
	fmt.Println("End Main")
}
