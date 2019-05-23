package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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

// WaitForInterrupt _
func WaitForInterrupt() chan bool {
	signal.Ignore(os.Interrupt)
	signal.Ignore(syscall.SIGINT)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	wait := make(chan bool, 1)

	go func() {
		for range sigc {
			wait <- true
		}
	}()

	return wait

}
