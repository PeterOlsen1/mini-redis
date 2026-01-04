package server

import "mini-redis/resp"

func HandlePing(args []resp.RESPItem) (string, error) {
	if len(args) == 0 {
		return "PONG", nil
	}

	return args[0].Content, nil
}
