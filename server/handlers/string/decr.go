package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
)

func HandleDecr(args []resp.RESPItem) (string, error) {
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
