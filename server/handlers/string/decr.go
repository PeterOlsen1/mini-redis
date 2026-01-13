package string

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleDecr(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.DECR, authtypes.WRITE)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.DECR, 1)
	}

	key := args.String(0)
	newVal, ok := internal.Decr(key)
	if !ok {
		return nil, errors.NOT_INTEGER
	}

	return resp.BYTE_STRING(newVal), nil
}
