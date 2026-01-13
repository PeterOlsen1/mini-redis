package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
)

func HandlePing(_ *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) == 0 {
		return resp.BYTE_STRING("PONG"), nil
	}

	return resp.BYTE_STRING(args[0].Content), nil
}
