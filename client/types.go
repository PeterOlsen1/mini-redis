package client

import (
	"net"
)

type RedisClient struct {
	addr string
	conn net.Conn
}

type ClientOptions struct {
	// Does not need to be filled out if URL is complete
	Addr string

	// redis://user:password@host:port/dbIdx (dbIdx is optional)
	URL string
}

type RequestBuilder struct {
	req     string
	len     int
	bufSize int
}
