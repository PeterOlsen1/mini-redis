package string

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleIncr(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.INCR, 1)
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.INCR, "WRITE", key)
	}
	newVal, ok := user.DB.Incr(key)
	if !ok {
		return nil, errors.NOT_INTEGER
	}

	return resp.BYTE_STRING(newVal), nil
}
