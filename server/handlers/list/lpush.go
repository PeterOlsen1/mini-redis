package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleLPush(user auth.User, args []resp.RESPItem) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.LPUSH, auth.WRITE)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.LPUSH, 2)
	}

	key := args[0].Content
	vals := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		vals[i] = arg.Content
	}

	ret := internal.LPush(key, vals)
	if ret == -1 {
		return nil, errors.WRONGTYPE
	}

	return resp.BYTE_INT(ret), nil
}
