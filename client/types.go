package client

import (
	"net"
)

type RedisClient struct {
	addr string
	conn net.Conn
}

type Options struct {
	Addr string
}

type RequestBuilder struct {
	req string
	len int
}
