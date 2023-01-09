package main

import (
	"fmt"
	"module/monolib"
)

func main(){
	fmt.Println("Hello World")
	err := monolib.TestListen()
	if err != nil {
		fmt.Println(err)
	}
}