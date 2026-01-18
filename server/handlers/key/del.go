package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleDel(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.DEL, 1)
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.DEL, authtypes.WRITE, key)
	}
	internal.Del(key)
	return resp.BYTE_OK, nil
}
