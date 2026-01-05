package server

import "mini-redis/resp"

func HandlePing(args []resp.RESPItem) ([]byte, error) {
	if len(args) == 0 {
		return resp.BYTE_STRING("PONG"), nil
	}

	return resp.BYTE_STRING(args[0].Content), nil
}
