package main

import (
	"flag"
	"mini-redis/server/start"
)

func main() {
	configPath := flag.String("cfg", "~/.mini-redis/config.yaml", "Location of configuration file")
	flag.Parse()

	start.Start(*configPath)
}
