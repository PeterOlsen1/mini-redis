package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleExpireTime(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.EXPIRETIME, 1)
	}

	key := args.String(0)
	if !user.CanRead(key) {
		return nil, errors.PERMS_KEY(commands.EXPIRETIME, "READ", key)
	}
	ret := user.DB.HandleExpireTime(key)

	return resp.BYTE_INT(int(ret)), nil
}
