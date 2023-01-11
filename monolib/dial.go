package monolib

import (
	"fmt"
	"net"
	"io"
)

func Dial() error {
	// 랜덤 포트에 리스너 생성
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer listener.Close()

	// make 는 채널을 생성합니다.
	// chan 은 채널 타입을 나타냅니다.

	// done 채널을 생성합니다.
	// 빈 구조체는 메모리를 사용하지 않습니다.
	done := make(chan struct{})

	// 리스너의 고루틴을 루프에서 처리하고 각 연결 처리 로직을 담당하는 것을 핸들러라고 합니다.
	go func() {

		// done
		defer func () { 
			fmt.Println("done")
			done <- struct{}{} 
		}()

		for {
			// 리스너에서 연결을 받습니다.
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}

			go func(c net.Conn) {
				defer func() {
					fmt.Println("conn done")
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					// 연결에서 데이터를 읽습니다.
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							fmt.Println(err)
						}
						return
					}
					fmt.Println("received",string(buf[:n]))
				}
			}(conn)
		}

	}()

	// dial은 name network에 대한 연결을 시도합니다.
	// network는 "tcp", "tcp4", "tcp6", "unix" or "unixpacket" 중 하나입니다.
	// name은 network에 따라 다릅니다.
	// 예를 들어, "tcp" network는 "host:port" 형식의 name을 받습니다.
	// "unix" network는 파일 시스템 경로를 받습니다.
	// "unixpacket" network는 "host:port" 형식의 name을 받습니다.
	// "host"는 IP 주소, IPv6 주소, 또는 호스트 이름이 될 수 있습니다.
	// "port"는 포트 번호 또는 서비스 이름이 될 수 있습니다.
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		fmt.Println("dial error", err)
		return err
	}

	conn.Close()
	<-done
	listener.Close()
	<-done
	return nil
}