package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types"
	"mini-redis/types/errors"
)

func HandleGet(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("get requires 1 argument")
	}

	key := args[0].Content
	val := internal.Get(key)
	if val == nil {
		return resp.BYTE_NULL, nil
	}

	if val.Type == types.STRING {
		strVal, ok := val.Item.(string)
		if !ok {
			return nil, fmt.Errorf(errors.WRONGTYPE)
		}
		return resp.BYTE_STRING(strVal), nil
	}

	return nil, fmt.Errorf(errors.WRONGTYPE)
}
