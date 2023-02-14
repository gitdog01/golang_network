// 여러개의 요청을 하나만 받고 나머지는 모두 취소하는 함수
// 예를 들어 여러개의 서버에서 TCP를 통해 단 하나의 리소스만 받아 올 필요가 있을 때 사용
package monolib

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

func TestDialContextCancelFanOut(){
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	// 해당 context 의 Err 매서드는 다음 중 하나를 반환합니다.
	// Canceled: context가 취소되었음을 나타냅니다.
	// DeadlineExceeded: context의 마감 시간이 지났음을 나타냅니다.
	// nil: context가 마감되지 않았음을 나타냅니다.

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go func(){
		// 하나의 연결만 수락
		conn, err := listener.Accept()
		if err == nil {
			conn.Close()
			// 연결이 수락되면 context를 취소
		}
	}()

	// 여러개의 연결을 시도를 추상화함 함수
	// 매개 변수로 주어진 주소로 연결을 시도하고
	// 연결이 성공하면 response 채널에 id를 보냄

	dial := func(ctx context.Context, address string, response chan int,id int, wg *sync.WaitGroup){
		defer wg.Done()

		var d net.Dialer
		c, err := d.DialContext(ctx, "tcp", address)
		if err != nil {
			return 
		}
		c.Close()

		select {
			case <-ctx.Done():
			case response <- id:
		}
	}

	res := make(chan int)
	var wg sync.WaitGroup

	// 10개의 연결을 시도
	// 10개의 다이얼러를 생성하는데
	// 다른 다이얼러가 먼저 연결되어 DialContext 함수의
	// 다이얼링이 블로킹된 경우 이를 해제하기 위해 cancel 함수를 호출하거나
	// 데드라인을 통해 취소합니다.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dial(ctx, listener.Addr().String(), res, i+1, &wg)
	}

	// 성공 연결이 있다면 res 채널에서 받습니다.
	// 이후 context를 취소하고 다른 다이얼러들의 연결 시도를 종료합니다.
	response := <-res
	cancel()
	// 이 지점에 wait는 다른 다이얼러들의 연결 시도를 중단하고 고루틴이 종료될 때까지 블로킹 됩니다.
	wg.Wait()
	close(res)

	if ctx.Err() != context.Canceled {
		fmt.Printf("Dial #%s succeeded, want canceled context", ctx.Err())
	}
	fmt.Printf("Dial #%d succeeded, want canceled context", response)
}