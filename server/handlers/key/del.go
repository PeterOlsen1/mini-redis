package key

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
)

func HandleDel(args []types.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("del requires 1 argument")
	}

	key := args[0].Content
	internal.Del(key)
	return "OK", nil
}
