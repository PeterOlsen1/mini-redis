package string

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleDecr(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.DECR, 1)
	}

	key := args[0].Content
	newVal, ok := internal.Decr(key)
	if !ok {
		return nil, errors.NOT_INTEGER
	}

	return resp.BYTE_STRING(newVal), nil
}
