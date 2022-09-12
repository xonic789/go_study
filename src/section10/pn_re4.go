package main

import (
	"fmt"
	"os"
)

func fileOpen(fileName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}

}

func main() {
	fileOpen("undefined.txt")
	fmt.Println("End Main")
}
