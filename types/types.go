package types

import (
	"mini-redis/server/auth"
	"net"
)

type Connection struct {
	Conn net.Conn
	*auth.User
}
