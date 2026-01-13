package main

import (
	"context"
	"flag"
	"fmt"
	"mini-redis/server/cfg"
	"mini-redis/server/internal"
	"mini-redis/server/log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	Run()
}

func Run() {
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
		fmt.Printf("Error: %e\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	log.StartLogger(ctx)
	internal.StartTTLScan(ctx)
	StartServer(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		cancel()
	}()
}
