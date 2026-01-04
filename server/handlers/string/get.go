package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types"
	"mini-redis/types/errors"
)

func HandleGet(args []resp.RESPItem) (string, error) {
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

	return "", fmt.Errorf(errors.WRONGTYPE)
}
