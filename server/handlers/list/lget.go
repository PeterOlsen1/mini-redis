package list

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleLGet(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.LGET, 1)
	}

	key := args[0].Content
	arr, err := internal.LGet(key)
	if err != nil {
		return nil, err
	}

	if arr == nil {
		return resp.BYTE_NULL, nil
	}

	serialized, err := resp.Serialize(arr, resp.ARRAY)
	if err != nil {
		return nil, err
	}

	return serialized, err
}
