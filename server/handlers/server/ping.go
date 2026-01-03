package server

import (
	"mini-redis/types"
)

func HandlePing(args []types.RESPItem) (string, error) {
	if len(args) == 0 {
		return "PONG", nil
	}

	return args[0].Content, nil
}
