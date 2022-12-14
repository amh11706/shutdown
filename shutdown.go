package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var tasks = make([]func(context.Context), 0)
var quit = make(chan os.Signal, 1)
var Closing bool

func init() {
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
}

func AddTask(f func(context.Context)) {
	if f != nil {
		tasks = append(tasks, f)
	}
}

// Trigger manually triggers a shutdown.
func Trigger() {
	log.Println("Shutting down")
	Closing = true
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	for _, t := range tasks {
		t(ctx)
	}
	cancel()
	log.Println("Shutdown complete")
}

// Watch listens to os.Interrupt or syscall.SIGTERM and triggers a shutdown.
func Watch() {
	<-quit
	Trigger()
}
