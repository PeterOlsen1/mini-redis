package string

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleGet(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.GET, 1)
	}

	key := args[0].Content
	val := internal.Get(key)
	if val == nil {
		return resp.BYTE_NULL, nil
	}

	if val.Type == types.STRING {
		strVal, ok := val.Item.(string)
		if !ok {
			return nil, errors.WRONGTYPE
		}
		return resp.BYTE_STRING(strVal), nil
	}

	return nil, errors.WRONGTYPE
}
