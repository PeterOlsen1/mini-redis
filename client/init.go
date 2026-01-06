package client

import (
	"net"
)

func NewClient(opt *ClientOptions) (*RedisClient, error) {
	var c RedisClient
	if opt == nil {
		opt = &ClientOptions{
			Addr: "localhost:6379",
		}
	}
	if opt.Addr == "" {
		c.addr = "localhost:6379"
	} else {
		c.addr = opt.Addr
	}

	// establish connection to the redis server
	if err := c.establishConnection(); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *RedisClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Establishes connection and sets connection in RedisClient
func (c *RedisClient) establishConnection() error {
	if c.conn != nil {
		return nil
	}

	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}
