package monitoring

import (
	"time"

	"github.com/trusz/idly/src/tracking/info"
)

// StartActiveAppMonitoring _
func StartActiveAppMonitoring(eventChannel *chan Event) {

	var lastActiveApp = info.ActiveApp()

	for true {
		var activeApp = info.ActiveApp()
		if activeApp != lastActiveApp {
			*eventChannel <- newAppEvent(activeApp)
		}
		lastActiveApp = activeApp
		time.Sleep(100 * time.Millisecond)
	}

}

// AppEvent _
type AppEvent struct {
	Event
	App string
}

func newAppEvent(appName string) AppEvent {
	return AppEvent{
		App: appName,
	}
}
