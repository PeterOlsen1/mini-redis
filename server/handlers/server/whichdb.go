package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
)

func HandleWhichDB(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	return resp.BYTE_INT(user.DB.Number), nil
}
