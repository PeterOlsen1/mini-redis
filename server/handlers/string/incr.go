package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/errors"
)

func HandleIncr(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("incr requires 1 argument")
	}

	key := args[0].Content
	newVal, ok := internal.Incr(key)
	if !ok {
		return nil, fmt.Errorf(errors.NOT_INTEGER)
	}

	return resp.BYTE_STRING(newVal), nil
}
