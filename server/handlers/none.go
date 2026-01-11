package handlers

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
)

func HandleNone(_ auth.User, args []resp.RESPItem) ([]byte, error) {
	return nil, fmt.Errorf("no command provided")
}
