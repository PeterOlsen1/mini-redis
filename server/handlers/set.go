package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/server/types"
)

func handleSet(args []types.RESPItem) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("set requires 2 arguments")
	}

	key := args[0].Content
	value := args[1].Content

	internal.Redis[key] = value
	return "OK", nil
}
