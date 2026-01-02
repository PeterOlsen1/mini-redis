package handlers

import (
	"fmt"
	"mini-redis/server/types"
)

func handleNone(args []types.RESPItem) (string, error) {
	return "", fmt.Errorf("no command provided")
}
