package handlers

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
)

func HandleNone(_ *authtypes.User, args resp.ArgList) ([]byte, error) {
	return nil, fmt.Errorf("no command provided")
}
