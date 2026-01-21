package string

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleDecr(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.DECR, 1)
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.DECR, "WRITE", key)
	}

	newVal, ok := user.DB.Decr(key)
	if !ok {
		return nil, errors.NOT_INTEGER
	}

	return resp.BYTE_STRING(newVal), nil
}
