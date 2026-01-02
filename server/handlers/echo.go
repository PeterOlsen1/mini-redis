package handlers

import (
	"mini-redis/types"
)

func handleEcho(args []types.RESPItem) (string, error) {
	if len(args) == 0 {
		return "ERROR ERROR!!!!!!", nil
	} else {
		return args[0].Content, nil
	}
}
