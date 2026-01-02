package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/server/types"
)

func handleGet(args []types.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("get requires 1 argument")
	}

	key := args[0].Content
	return internal.Get(key), nil
}
