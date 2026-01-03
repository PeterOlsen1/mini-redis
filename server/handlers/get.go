package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
	"strconv"
)

func handleGet(args []types.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("get requires 1 argument")
	}

	key := args[0].Content
	val := internal.Get(key)
	if val == nil {
		return "", nil
	}

	if val.Type == types.STRING {
		return val.Item.(string), nil
	}

	if val.Type == types.INT {
		return strconv.Itoa(val.Item.(int)), nil
	}

	return "", fmt.Errorf("Operation against a key holding the wrong kind of value")
}
