package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/errors"
)

func HandleDecr(args []resp.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("decr requires 1 argument")
	}

	key := args[0].Content
	newVal, ok := internal.Decr(key)
	if !ok {
		return "", fmt.Errorf(errors.NOT_INTEGER)
	}

	return newVal, nil
}
