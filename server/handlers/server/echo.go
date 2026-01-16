package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleEcho(_ *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.ECHO, 1)
	} else {
		return resp.BYTE_STRING(args.String(0)), nil
	}
}
