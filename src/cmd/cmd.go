package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

// Run _
func Run(command string) string {

	shell := "sh"
	arg0 := "-c"

	out, err := exec.Command(shell, arg0, command).Output()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", out)
}
