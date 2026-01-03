package client

import (
	"fmt"
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
}

func InitRequest(arrLen int, command string) *RequestBuilder {
	return &RequestBuilder{
		req: fmt.Sprintf("*%d\r\n$%d\r\n%s\r\n", arrLen, len(command), command),
	}
}

func (r *RequestBuilder) AddParam(param string) *RequestBuilder {
	r.req += fmt.Sprintf("$%d\r\n%s\r\n", len(param), param)
	return r
}
