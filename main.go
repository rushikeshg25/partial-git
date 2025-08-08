package main

import (
	"context"
	"os"
	"os/signal"
	"partial-git/cmd"
	"syscall"
)

var Version = "PLACEHOLDER"

func run() int {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
		<-c
		os.Exit(1)
	}()

	if err := cmd.Execute(ctx); err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
