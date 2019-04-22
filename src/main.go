package main

import (
	"log"

	"github.com/trusz/idly/src/info"
)

func main() {
	log.Println("Hello, World!")
	info.ActiveApp()
	info.IdleTime()
}
