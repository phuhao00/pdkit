package ksignal

import (
	"os"
	"os/signal"
	"syscall"
)

// InitSignal register signals handler.
func InitSignal() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	return c
}

// HandleSignal fetch signal from chan then do exit or reload.
func HandleSignal(c chan os.Signal, hfs map[syscall.Signal]func() error) func()error{
	// Block until a signal is received.
	hf := func(sig syscall.Signal) error {
		if f,exist:= hfs[syscall.SIGQUIT];exist {
			return f()
		}
		return nil
	}
	return func() error {
		for {
			s := <-c
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM,syscall.SIGINT:
				return hf(syscall.SIGQUIT)
			case syscall.SIGHUP:
				// TODO reload
				return hf(syscall.SIGQUIT)
			default:
				return hf(syscall.SIGQUIT)
			}
		}
	}
}