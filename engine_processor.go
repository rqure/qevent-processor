package main

import (
	"os"
	"os/signal"
	"syscall"

	qmq "github.com/rqure/qmq/src"
)

type EngineProcessor struct {
}

func (e *EngineProcessor) Process(cp qmq.EngineComponentProvider) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			return
		}
	}
}
