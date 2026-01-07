package main

import (
	"context"
	"flag"
	"fmt"
	"mini-redis/server/cfg"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	configPath := flag.String("cfg", "~/.mini-redis/config.yaml", "Location of configuration file")
	flag.Parse()

	if (*configPath)[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to get home directory:", err)
			os.Exit(1)
		}
		*configPath = filepath.Join(homeDir, (*configPath)[1:])
	}

	err := cfg.LoadConfig(*configPath)
	if err != nil {
		fmt.Println("Failed to read config, exiting")
		os.Exit(1)
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
