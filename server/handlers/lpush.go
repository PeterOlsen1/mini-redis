package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
)

func handleLPush(args []types.RESPItem) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("LPUSH requires 2 arguments")
	}

	key := args[0].Content
	newVal, ok := internal.Incr(key)
	if !ok {
		return "", fmt.Errorf("value is not an integer or out of range")
	}

	return newVal, nil
}
