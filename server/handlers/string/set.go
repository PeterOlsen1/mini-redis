package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleSet(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("set requires 2 arguments")
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.SET, authtypes.WRITE, key)
	}
	value := args.String(1)

	internal.Set(key, value)
	return resp.BYTE_OK, nil
}
