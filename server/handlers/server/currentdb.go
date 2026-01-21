package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"strconv"
)

func HandleWhichDB(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	return resp.BYTE_STRING(strconv.Itoa(user.DB.Number)), nil
}
