package handlers

import (
	"fmt"
	"mini-redis/resp"
)

func HandleNone(args []resp.RESPItem) (string, error) {
	return "", fmt.Errorf("no command provided")
}
