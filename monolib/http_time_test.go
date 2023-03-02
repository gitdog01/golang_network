package monolib

// HTTP 기본 워크 플로우
import (
	"fmt"
	"net/http"
	"time"
)

func TestHeadTime(){
	resp, err := http.Head("http://www.time.gov/")
	if err != nil {
		panic(err)
	}

	_ = resp.Body.Close() // 예외 상황 처리 없이 바로 받음
	// Body 내용을 보지 않더라도 close 를 해야한다.
	// Body 를 읽지 않으면 커넥션을 계속 유지하고 있어서 ( keepalive )
	// 커넥션 풀이 꽉 차게 된다. 

	now := time.Now().Round(time.Second)
	data := resp.Header.Get("Date")
	if data == "" {
		panic("no date")
	}

	dt, err := time.Parse(time.RFC1123, data)
	// RFC1123은 HTTP 헤더에 사용되는 날짜 형식입니다.
	if err != nil {
		panic(err)
	}

	fmt.Println("now:", now)
	fmt.Println("data:", dt)

}