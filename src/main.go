package main

import (
	"github.com/trusz/idly/src/logger"
	"github.com/trusz/idly/src/monitoring"

	"github.com/trusz/idly/src/cmd"
)

func main() {
	var eventChannel = make(chan monitoring.Event)

	go func() {
		monitoring.StartActiveAppMonitoring(&eventChannel)
	}()

	go func() {
		monitoring.StartIdleMonitoring(&eventChannel)
	}()

	go func() {
		logger.StartLogger(&eventChannel)
	}()

	<-cmd.WaitForInterrupt()

}
