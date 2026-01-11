package string

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleGet(user *auth.User, args resp.ArgList) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.GET, auth.READ)
	}

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
