package types

import (
	"mini-redis/server/auth/authtypes"
	"net"
)

type Connection struct {
	Conn net.Conn
	*authtypes.User
}
