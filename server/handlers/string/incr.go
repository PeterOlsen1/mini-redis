package string

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleIncr(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.INCR, authtypes.WRITE)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.INCR, 1)
	}

	key := args.String(0)
	newVal, ok := internal.Incr(key)
	if !ok {
		return nil, errors.NOT_INTEGER
	}

	return resp.BYTE_STRING(newVal), nil
}
