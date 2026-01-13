package list

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strconv"
)

func HandleRPop(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.RPOP, authtypes.WRITE)
	}

	if len(args) < 1 {
		return nil, fmt.Errorf("ERR wrong number of arguments for 'rpop' command")
	}

	key := args.String(0)
	if len(args) == 1 {
		res, err := internal.RPop(key, 1)
		if err != nil {
			return nil, err
		}

		if len(res) == 0 {
			return nil, fmt.Errorf("cannot pop empty list")
		}

		return resp.BYTE_STRING(res[0]), nil
	}

	num, err := strconv.Atoi(args[1].Content)
	if err != nil {
		return nil, fmt.Errorf("invalid pop count")
	}
	res, err := internal.RPop(key, num)
	if err != nil {
		return nil, err
	}

	serialized, err := resp.Serialize(res, resp.ARRAY)
	if err != nil {
		return nil, err
	}

	return serialized, nil
}
