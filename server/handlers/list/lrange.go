package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strconv"
)

func HandleLRange(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.LRANGE, authtypes.READ)
	}

	if len(args) < 3 {
		return nil, errors.ARG_COUNT(commands.LRANGE, 3)
	}

	// parse arguments
	key := args.String(0)
	start := args[1].Content
	end := args[2].Content
	startInt, err := strconv.Atoi(start)
	if err != nil {
		return nil, errors.INVALID_ARG
	}
	endInt, err := strconv.Atoi(end)
	if err != nil {
		return nil, errors.INVALID_ARG
	}

	internalResp, err := internal.LRange(key, startInt, endInt)
	if err != nil {
		return nil, err
	}

	serialized, err := resp.Serialize(internalResp, resp.ARRAY)
	if err != nil {
		return nil, err
	}

	return serialized, nil
}
