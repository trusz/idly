package info

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trusz/idly/src/tracking/cmd"
)

// ActiveApp _
func ActiveApp() string {
	const command = "lsappinfo info -only name $(lsappinfo front)"
	var output = cmd.Run(command)

	var infos = strings.Split(output, "=")

	if len(infos) < 2 {
		return "NONE"
	}

	var quatedAppName = infos[1]
	var appName = strings.Replace(quatedAppName, "\"", "", -1)
	appName = strings.TrimSuffix(appName, "\n")

	return appName
}

// IdleSeconds _
func IdleSeconds() int {
	const command = "/usr/sbin/ioreg -c IOHIDSystem | /usr/bin/awk '/HIDIdleTime/ {print int($NF/1000000000); exit}'"
	var output = cmd.Run(command)
	output = strings.TrimSuffix(output, "\n")

	idle, err := strconv.Atoi(output)
	if err != nil {
		fmt.Println(err)
	}

	return idle
}
