// 127.0.0.1 주소에서 랜덤 포트에 수신 대기 중인 리스너 생성
package monolib

import (
	"net"
	"fmt"
)

func TestListen() error {
	// listener 와 err 를 리턴 받습니다.
	// listener 는 net.Listener 인터페이스를 구현한 객체입니다.
	// err 는 error 인터페이스를 구현한 객체입니다.
	// net.Listen 은 net 패키지의 Listen 함수를 호출합니다.
	// Listen 함수는 "tcp" 네트워크 프로토콜을 사용합니다.
	// 포트를 표기하거나 0으로 하면 무작위 포트를 사용합니다.

	// := 는 변수 선언과 초기화를 동시에 합니다.
	// := 는 변수를 선언하고 초기화하는 것이기 때문에
	// 변수를 선언하고 초기화하는 것이 아니라면 := 를 사용할 수 없습니다.
	// := 는 함수 내에서만 사용할 수 있습니다.

	listener, err := net.Listen("tcp", "127.0.0.1:0")

	// 에러가 발생하면 테스트를 실패합니다.
	// 에러가 발생하지 않으면 테스트를 성공합니다.
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// defer 는 함수가 종료되기 직전에 실행됩니다.
	// defer 은 역순으로 진행됩니다.
	// defer 는 함수 내에서만 사용할 수 있습니다.
	// defer 는 함수가 종료되기 직전에 실행되기 때문에
	// 리소스를 해제하는 용도로 사용합니다.
	// 명시적으로 리소스를 해제하지 않는다면 메모리 누수가 발생하거나 
	// Accept 메서드가 무한정 블로킹되는 데드락이 발생할 수 있습니다.

	// _ 는 무시할 값을 의미합니다.
	defer func() { _ = listener.Close() }()

	// listener.Addr()의 반환값은 
	// %q 는 쌍따옴표로 묶인 문자열을 표시합니다.
	fmt.Println("bound to", listener.Addr())

	// for 루프는 무한 루프입니다.

	// Accept 메서드는 리스너가 수신 대기 중인 연결을 반환합니다.
	// Accept 메서드는 블로킹 메서드입니다.
	// Accept 메서드는 블로킹 메서드이기 때문에
	// 리스너가 수신 대기 중인 연결이 없으면 무한정 블로킹됩니다.
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		// go 는 고루틴을 생성합니다.
		// 고루틴은 경량스레드 입니다.
		go func (c net.Conn) {
			defer c.Close()
			_, _ = c.Write([]byte("Hello World"))

		}(conn)
	}
}