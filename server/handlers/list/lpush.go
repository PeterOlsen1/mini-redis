package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleLPush(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.LPUSH, 2)
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.LPUSH, "WRITE", key)
	}

	vals := args.Slice(1, 10000)

	ret := user.DB.LPush(key, vals)
	if ret == -1 {
		return nil, errors.WRONGTYPE
	}

	return resp.BYTE_INT(ret), nil
}
