package string

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleGet(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.GET, 1)
	}

	key := args.String(0)
	if !user.CanRead(key) {
		return nil, errors.PERMS_KEY(commands.GET, "READ", key)
	}

	val := user.DB.Get(key)
	if val == nil {
		return resp.BYTE_NULL, nil
	}

	if val.Type == internal.STRING {
		strVal, ok := val.Item.(string)
		if !ok {
			return nil, errors.WRONGTYPE
		}
		return resp.BYTE_STRING(strVal), nil
	}

	return nil, errors.WRONGTYPE
}
