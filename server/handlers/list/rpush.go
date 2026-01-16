package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleRPush(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.RPUSH, authtypes.WRITE)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.RPUSH, 2)
	}

	key := args.String(0)
	vals := args.Slice(1, 10000)

	ret := internal.RPush(key, vals)
	if ret == -1 {
		return nil, errors.WRONGTYPE
	}

	return resp.BYTE_INT(ret), nil
}
