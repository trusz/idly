package monitoring

import (
	"time"

	"github.com/trusz/idly/src/tracking/info"
)

// StartIdleMonitoring _
func StartIdleMonitoring(eventChannel *chan Event) {
	const consideredIdleSeconds = 10 * 60
	var isIdle = false

	for true {
		var idleSeconds = info.IdleSeconds()
		if !isIdle && idleSeconds >= consideredIdleSeconds {
			isIdle = true
			*eventChannel <- newIdleEvent(isIdle)
		}
		if isIdle && idleSeconds < consideredIdleSeconds {
			isIdle = false
			*eventChannel <- newIdleEvent(isIdle)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

// IdleEvent _
type IdleEvent struct {
	Event
	IsIdle bool
}

func newIdleEvent(isIdle bool) IdleEvent {
	return IdleEvent{
		IsIdle: isIdle,
	}
}
