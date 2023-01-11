package monolib

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

// DialTimeout 함수는 Dial 함수와 동일하게 동작하지만
// 지정된 시간이 지나면 연결을 종료합니다.
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	// net.Dialer 인터페이스에 대한 제어권을 제공하지 않기 때문에
	// net.Dialer 인터페이스를 사용하여 커스텀 Dialer를 만듭니다.
	// DialTimeout 함수는 에러를 반환하기 위한 net.Dialer 인터페이스의 Control 함수를 오버라이딩합니다.
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TestDialTimeout() {
	// 5초 안에 성공하지 못하면 타임아웃 됩니다.
	// 이 예시에서는 10.0.0.1 이라는 연결할 수 없는 주소를 사용합니다.
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err != nil {
		c.Close()
		fmt.Println(err)
	}

	// 타임 아웃 메서드에서 확인하기 이전에 먼저 에러를 net.Error로 타입 어설션합니다.
	nErr, ok := err.(net.Error)
	if !ok {
		fmt.Println("not a net.Error")
	}
	if !nErr.Timeout() {
		fmt.Println("not a timeout")
	}
}
// 여러 IP 주소로 해석되는 호스트에 다이얼 시 Go 에서는 각 IP 주소 중
// 먼저 연결되는 주소를 기본 IP 주소로 연결을 시도합니다.
// 첫 번째 연결이 성공하면 나머지 연결은 취소됩니다.
// 모든 연결이 실패하면 마지막 연결의 에러가 반환됩니다.