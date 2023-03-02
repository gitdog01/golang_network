// 하트비트 구현하기
package monolib

import (
	"context"
	"io"
	"time"
)

const defaultPingInterval = 30 * time.Second

func pringer(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
		case <-ctx.Done():
			return
		case interval = <-reset:
			default:
	}
	if interval <= 0 {
		interval = defaultPingInterval
	}

	timer := time.NewTimer(interval)
	defer func(){
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
			case <-ctx.Done():
				return
			case interval = <-reset:
				if interval <= 0 {
					interval = defaultPingInterval
				}
				if !timer.Stop() {
					<-timer.C
				}
			case <-timer.C:
		}
		if _, err := w.Write([]byte("ping")); err != nil {
			return
		}
		_ = timer.Reset(interval)
	}
}