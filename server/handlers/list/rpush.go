package list

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleRPush(user *auth.User, args resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.RPUSH, auth.WRITE)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.RPUSH, 2)
	}

	key := args.String(0)
	vals := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		vals[i] = arg.Content
	}

	ret := internal.RPush(key, vals)
	if ret == -1 {
		return nil, errors.WRONGTYPE
	}

	return resp.BYTE_INT(ret), nil
}
