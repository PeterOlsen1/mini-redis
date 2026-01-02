package handlers

import (
	"mini-redis/server/types"
)

func handlePing(args []types.RESPItem) (string, error) {
	if len(args) == 0 {
		return "PONG", nil
	}

	return args[0].Content, nil
}
