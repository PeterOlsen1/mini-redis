package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/info"
	"mini-redis/server/internal"
	"strconv"
)

func HandleInfo(_ *auth.User, _ resp.ArgList) ([]byte, error) {
	// take information given from server, add total keys as well
	respInfo := info.GetInfo()
	respInfo += "-Total keys: " + strconv.Itoa(internal.GetStoreSize()) + "\n"
	return resp.BYTE_STRING(respInfo), nil
}
