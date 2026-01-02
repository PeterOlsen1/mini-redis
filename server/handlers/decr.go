package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
)

func handleDecr(args []types.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("decr requires 1 argument")
	}

	key := args[0].Content
	newVal, ok := internal.Decr(key)
	if !ok {
		return "", fmt.Errorf("value is not an integer or out of range")
	}

	return newVal, nil
}
