package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleEcho(_ *auth.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.ECHO, 1)
	} else {
		return resp.BYTE_STRING(args[0].Content), nil
	}
}
