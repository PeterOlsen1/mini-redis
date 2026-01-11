package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
)

func HandleEcho(_ auth.User, args []resp.RESPItem) ([]byte, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Echo requires 1 argument")
	} else {
		return resp.BYTE_STRING(args[0].Content), nil
	}
}
