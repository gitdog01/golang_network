// dial_timeout 과 같지만
// context 를 사용하여 timeout 을 설정할 수 있다.
// context 는 비동기 프로세스에 취소 시그널을 보낼 수 있는 객체

// 콘텍스트로 취소하기 위해서는 초기화시 반환되는 cancel 함수를 호출해야 한다.
// cancel 함수를 코드상의 다른 부분으로 제어권을 넘길 수도 있습니다.

// 예를 들어, 사용자로 부터 Ctrl-C 입력받는 것처럼 운영체제의 특정 시그널을 모니터링하여
// 프로그램이 종료하기 전에 연결 시도를 우아하게 종단하고 이미 존재하는 연결을 제거할 수 있습니다.
// TO DO : 예전에 PM2 와 node 우아하게 종료한 글이랑 합치기

package monolib

import (
	"context"
	"fmt"
	"net"
	"syscall"
	"time"
)

func TestDialContext() {
	// 5초 안에 성공하지 못하면 타임아웃 됩니다.
	// 이 예시에서는
	dl := time.Now().Add(5 * time.Second)
	// 데드라인을 만듭니다.
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	// 사용후 바로 정리될 수 있도록 정의
	defer cancel()
	

	var d net.Dialer // DialContext의 메서드입니다.
	// Control 함수를 오버라이딩하여 여녈을 콘텍스트에 간신히 초과하는 정도로 지연시킵니다.
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(5 * time.Second + time.Millisecond)
		return nil
	}

	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err != nil {
		conn.Close()
		fmt.Println("err")
	}

	nErr, ok := err.(net.Error)
	if !ok {
		fmt.Println("not a net.Error")
	} else {
		if !nErr.Timeout() {
			fmt.Println("not a timeout")
		}
	}

	// 데드라인이 콘텍스트를 제대로 취소하였는지, Cancel 함수 호출에 문제는 없었는지를 확인합니다.
	if ctx.Err() != context.DeadlineExceeded {
		fmt.Println("not a deadline exceeded")
	}

}