package string

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleSet(user *auth.User, args []resp.RESPItem) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.SET, auth.WRITE)
	}

	if len(args) < 2 {
		return nil, fmt.Errorf("set requires 2 arguments")
	}

	key := args[0].Content
	value := args[1].Content

	internal.Set(key, value)
	return resp.BYTE_OK, nil
}
