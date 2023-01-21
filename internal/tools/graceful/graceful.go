package graceful

import (
	"context"
	"fmt"
	"limiter/internal/tools/closer"
	"os"
	"os/signal"
	"syscall"
)

var called bool

func Graceful() context.Context {
	if called {
		panic(fmt.Errorf("Graceful already called"))
	}

	called = true

	c := make(chan os.Signal, 1)

	ctx, cancel := context.WithCancel(context.Background())

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-c:
				cancel()
				closer.Closer.Close()
			default:
			}
		}
	}()

	return ctx
}
