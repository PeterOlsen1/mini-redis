package server

import (
	"fmt"
	"mini-redis/resp"
)

func HandleEcho(args []resp.RESPItem) ([]byte, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Echo requires 1 argument")
	} else {
		return resp.BYTE_STRING(args[0].Content), nil
	}
}
