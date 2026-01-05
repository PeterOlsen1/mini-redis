package list

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"strconv"
)

func HandleRPop(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("ERR wrong number of arguments for 'rpop' command")
	}

	key := args[0].Content
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
