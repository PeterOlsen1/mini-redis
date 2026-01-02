package handlers

import (
	"fmt"
	"mini-redis/types"
)

func handleNone(args []types.RESPItem) (string, error) {
	return "", fmt.Errorf("no command provided")
}
