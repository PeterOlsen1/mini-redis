package handlers

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
)

func HandleNone(_ *auth.User, args resp.ArgList) ([]byte, error) {
	return nil, fmt.Errorf("no command provided")
}
