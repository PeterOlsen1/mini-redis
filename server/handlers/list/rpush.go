package list

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/errors"
)

func HandleRPush(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("RPUSH requires 2 arguments")
	}

	key := args[0].Content
	vals := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		vals[i] = arg.Content
	}

	ret := internal.RPush(key, vals)
	if ret == -1 {
		return nil, fmt.Errorf(errors.WRONGTYPE)
	}

	return resp.BYTE_INT(ret), nil
}
