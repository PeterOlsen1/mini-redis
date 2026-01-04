package server

import "mini-redis/resp"

func HandleEcho(args []resp.RESPItem) (string, error) {
	if len(args) == 0 {
		return "ERROR ERROR!!!!!!", nil
	} else {
		return args[0].Content, nil
	}
}
