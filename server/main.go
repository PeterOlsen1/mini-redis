package main

import (
	"context"
	"fmt"
	"mini-redis/server/cfg"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if cfg.InitConfig() != nil {
		fmt.Println("Failed to read configuration")
		os.Exit(-1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	StartServer(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		cancel()
	}()
}
