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
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.SET, authtypes.WRITE)
	}

	if len(args) < 2 {
		return nil, fmt.Errorf("set requires 2 arguments")
	}

	key := args.String(0)
	value := args[1].Content

	internal.Set(key, value)
	return resp.BYTE_OK, nil
}
