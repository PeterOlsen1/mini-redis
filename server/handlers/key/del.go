package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleDel(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.DEL, authtypes.WRITE)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.DEL, 1)
	}

	key := args.String(0)
	internal.Del(key)
	return resp.BYTE_OK, nil
}
