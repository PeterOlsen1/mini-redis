package handlers

import (
	"fmt"
	"mini-redis/types"
)

func HandleNone(args []types.RESPItem) (string, error) {
	return "", fmt.Errorf("no command provided")
}
