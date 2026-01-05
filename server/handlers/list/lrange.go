package list

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/errors"
	"strconv"
)

func HandleLRange(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("LRANGE requires 3 arguments")
	}

	// parse arguments
	key := args[0].Content
	start := args[1].Content
	end := args[2].Content
	startInt, err := strconv.Atoi(start)
	if err != nil {
		return nil, fmt.Errorf(errors.INVALID_ARG)
	}
	endInt, err := strconv.Atoi(end)
	if err != nil {
		return nil, fmt.Errorf(errors.INVALID_ARG)
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
