// TCP 로 데이터가 올때마다 버퍼 크기 만큼 읽어서 처리한다.
package monolib

import (
	"crypto/rand"
	"fmt"
	"net"
)

func ReadIntoBuffer(){
	payload := make([]byte, 1<<24) // 16MB
	_, err := rand.Read(payload)  // 랜덤한 데이터를 payload 에 채운다.
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		_, err = conn.Write(payload) 
		if err != nil {
			panic(err)
		}
	}()

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 데이터가 올때 마다 버퍼 크기 만큼 읽어서 처리한다.

	buf := make([]byte, 1<<19) // 512KB

	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("Read %d bytes from the connection into the buffer \r", n)	
	}
}