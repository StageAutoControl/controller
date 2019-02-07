package internal

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// NewExitHandlerContext creates a trap for termination signals
func NewExitHandlerContext(logger *logrus.Logger) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		defer cancel()
		logger.Info("shutting down")
	}()

	return ctx
}
