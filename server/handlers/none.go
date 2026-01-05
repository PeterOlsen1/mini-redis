package handlers

import (
	"fmt"
	"mini-redis/resp"
)

func HandleNone(args []resp.RESPItem) ([]byte, error) {
	return nil, fmt.Errorf("no command provided")
}
