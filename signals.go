package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// Signals is a channel that intercepts Interrupt signals.
//
// The channel is global and public so that other packages
// may send Interrupt signals to initiate the stopFunc.
var Signals chan os.Signal

func init() {
	Signals = make(chan os.Signal, 1)
	signal.Notify(Signals, os.Interrupt, syscall.SIGTERM)
}

// Interrupt takes a stopFunc and executes it whenever the
// app receives an Interrupt signal.
//
// The stopFunc is typically configured to initiate a clean
// shutdown of the application.
//
// This function can be the last statement in func main()
// where the return parameters are discarded, but it can
// also be the last statement in a CLI action where the
// returned error is used to determine the exit code.
func Interrupt(stopFunc func() error) error {
	<-Signals
	return stopFunc()
}
