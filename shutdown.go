package shutdown

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var tasks = make([]func(), 0)
var quit = make(chan os.Signal, 1)
var Closing bool

func init() {
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
}

func AddTask(f func()) {
	if f != nil {
		tasks = append(tasks, f)
	}
}

func Watch() {
	_ = <-quit
	log.Println("Shutting down")
	Closing = true
	for _, t := range tasks {
		t()
	}
	log.Println("Shutdown complete")
}
