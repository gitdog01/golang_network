package main

import (
	"fmt"
	"os"
	"module/monolib"
)

func main(){

	// args를 가져옵니다.
	args := os.Args
	var listenerFlag bool = true

	// args가 2개보다 크면 에러를 출력하고 종료합니다.
	if len(args) > 2 {
		fmt.Println("Usage: go run main.go -c")
		return
	}

	if args[1] == "-c" {
		listenerFlag = false
	}	

	if listenerFlag {
		// 리스너 모드
		err := monolib.RunListen()
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}