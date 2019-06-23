package tracking

import (
	"github.com/trusz/idly/src/tracking/cmd"
	"github.com/trusz/idly/src/tracking/logger"
	"github.com/trusz/idly/src/tracking/monitoring"
	"github.com/trusz/idly/src/tracking/storage"
)

// Start _
func Start() {
	var eventChannel = make(chan monitoring.Event)
	var fileStorage = storage.New()

	go func() {
		monitoring.StartActiveAppMonitoring(&eventChannel)
	}()

	go func() {
		monitoring.StartIdleMonitoring(&eventChannel)
	}()

	go func() {
		logger.StartLogger(&eventChannel, fileStorage)
	}()

	<-cmd.WaitForInterrupt()
}
