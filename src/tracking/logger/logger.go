package logger

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/trusz/idly/src/tracking/monitoring"
)

// StartLogger _
func StartLogger(eventChannel *chan monitoring.Event, storage io.Writer) {
	for true {
		event := readFromChannel(eventChannel)
		logEntry := renderEvent(event)
		_, err := storage.Write([]byte(logEntry))
		if err != nil {
			log.Println(err)
		}
	}
}

func readFromChannel(eventChannel *chan monitoring.Event) monitoring.Event {
	event := <-*eventChannel
	return event
}

func renderEvent(event monitoring.Event) string {
	var logEntry string
	switch event.(type) {
	case monitoring.AppEvent:
		appEvent := event.(monitoring.AppEvent)
		logEntry = renderAppEvent(appEvent)
	case monitoring.IdleEvent:
		idleEvent := event.(monitoring.IdleEvent)
		logEntry = renderIdleEvent(idleEvent)
	default:
		logEntry = renderDefault()
	}

	currentTime := time.Now()
	formattedTimeStamp := currentTime.Format("2006.01.02 15:04:05")
	return formattedTimeStamp + " " + logEntry + "\n"

}

func renderAppEvent(event monitoring.AppEvent) string {
	return fmt.Sprintf("Switched to app: %s", event.App)
}

func renderIdleEvent(event monitoring.IdleEvent) string {
	return fmt.Sprintf("Is user idle: %t\n", event.IsIdle)
}

func renderDefault() string {
	return fmt.Sprintf("Unknown event")
}
