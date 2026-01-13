package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
)

func HandleWhoami(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if user == nil || user.Username == "" {
		return nil, fmt.Errorf("user is not authenticated")
	}

	return resp.BYTE_STRING(user.Username), nil
}
