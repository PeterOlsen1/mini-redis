package string

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
)

func HandleIncr(args []types.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("incr requires 1 argument")
	}

	key := args[0].Content
	newVal, ok := internal.Incr(key)
	if !ok {
		return "", fmt.Errorf("value is not an integer or out of range")
	}

	return newVal, nil
}
