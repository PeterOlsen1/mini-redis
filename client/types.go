package client

import (
	"net"
)

type RedisClient struct {
	addr string
	conn net.Conn
}

type ClientOptions struct {
	Addr string
}

type RequestBuilder struct {
	req     string
	len     int
	bufSize int
}
