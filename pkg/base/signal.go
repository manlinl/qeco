package base

import (
	"os"
	"os/signal"
	"syscall"
)

// SetupShutdownSignalHandler listens to SIGINT an SIGTERM. It closes `stopCh` to trigger graceful
// shutdown on first signal and exits on second one.
func SetupShutdownSignalHandler() <-chan struct{} {
	stopCh := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		close(stopCh)
		<-c
		os.Exit(1)
	}()
	return stopCh
}
