package logger

import (
	"log"

	"github.com/trusz/idly/src/monitoring"
)

// StartLogger _
func StartLogger(eventChannel *chan monitoring.Event) {
	for true {
		event := <-*eventChannel

		switch event.(type) {
		case monitoring.AppEvent:
			appEvent := event.(monitoring.AppEvent)
			log.Printf("Switched to app: %s", appEvent.App)
		case monitoring.IdleEvent:
			idleEvent := event.(monitoring.IdleEvent)
			log.Printf("Is user idle: %t\n", idleEvent.IsIdle)
		default:
			log.Println("Unknown event")
		}

	}
}
