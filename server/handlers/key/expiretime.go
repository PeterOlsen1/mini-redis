package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleExpireTime(user *auth.User, args resp.ArgList) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.EXPIRETIME, auth.READ)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.EXPIRETIME, 1)
	}

	key := args.String(0)
	ret := internal.HandleExpireTime(key)

	return resp.BYTE_INT(int(ret)), nil
}
