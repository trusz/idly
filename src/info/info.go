package info

import (
	"log"
	"strings"

	"github.com/trusz/idly/src/cmd"
)

// ActiveApp _
func ActiveApp() {
	const command = "lsappinfo info -only name $(lsappinfo front)"
	var output = cmd.Run(command)

	log.Printf("Active app:%s\n", output)
	var quatedAppName = strings.Split(output, "=")[1]
	var appName = strings.Replace(quatedAppName, "\"", "", -1)
	log.Printf("Active app: %s\n", appName)
}

// IdleTime _
func IdleTime() {
	const command = "/usr/sbin/ioreg -c IOHIDSystem | /usr/bin/awk '/HIDIdleTime/ {print int($NF/1000000000); exit}'"
	var output = cmd.Run(command)

	log.Printf("Idled for:%ss \n", output)
}
