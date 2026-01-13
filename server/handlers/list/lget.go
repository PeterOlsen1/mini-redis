package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleLGet(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.LGET, authtypes.READ)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.LGET, 1)
	}

	key := args.String(0)
	arr, err := internal.LGet(key)
	if err != nil {
		return nil, err
	}

	if arr == nil {
		return resp.BYTE_NULL, nil
	}

	serialized, err := resp.Serialize(arr, resp.ARRAY)
	if err != nil {
		return nil, err
	}

	return serialized, err
}
