package string

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
)

func HandleSet(args []types.RESPItem) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("set requires 2 arguments")
	}

	key := args[0].Content
	value := args[1].Content

	internal.Set(key, value)
	return "OK", nil
}
