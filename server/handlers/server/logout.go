package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
)

func HandleLogout(user *auth.User, args []resp.RESPItem) ([]byte, error) {
	if user == nil {
		return nil, fmt.Errorf("user is not authenticated")
	}

	user.Username = ""
	user.Perms = 0

	return resp.BYTE_OK, nil
}
