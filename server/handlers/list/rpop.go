package list

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleRPop(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("ERR wrong number of arguments for 'rpop' command")
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.RPOP, "WRITE", key)
	}

	if len(args) == 1 {
		res, err := user.DB.RPop(key, 1)
		if err != nil {
			return nil, err
		}

		if len(res) == 0 {
			return nil, fmt.Errorf("cannot pop empty list")
		}

		return resp.BYTE_STRING(res[0]), nil
	}

	num, err := args.Int(1)
	if err != nil {
		return nil, fmt.Errorf("invalid pop count")
	}
	res, err := user.DB.RPop(key, num)
	if err != nil {
		return nil, err
	}

	serialized, err := resp.Serialize(res, resp.ARRAY)
	if err != nil {
		return nil, err
	}

	return serialized, nil
}
