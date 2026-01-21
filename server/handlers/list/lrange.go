package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleLRange(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 3 {
		return nil, errors.ARG_COUNT(commands.LRANGE, 3)
	}

	// parse arguments
	key := args.String(0)
	if !user.CanRead(key) {
		return nil, errors.PERMS_KEY(commands.LRANGE, "READ", key)
	}

	start, startErr := args.Int(1)
	end, endErr := args.Int(2)
	if startErr != nil || endErr != nil {
		return nil, errors.INVALID_ARG
	}

	internalResp, err := user.DB.LRange(key, start, end)
	if err != nil {
		return nil, err
	}

	serialized, err := resp.Serialize(internalResp, resp.ARRAY)
	if err != nil {
		return nil, err
	}

	return serialized, nil
}
