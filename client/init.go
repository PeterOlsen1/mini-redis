package client

import (
	"fmt"
	"net"
	"strings"
)

func NewClient(opt *ClientOptions) (*RedisClient, error) {
	var c RedisClient
	if opt == nil {
		opt = &ClientOptions{
			Addr: "localhost:6379",
		}
	}

	if opt.URL != "" {
		trimmed := strings.TrimPrefix(opt.URL, "redis://")
		parts := strings.Split(trimmed, "@")

		if len(parts) != 2 {
			return nil, fmt.Errorf("failed to parse connection string")
		}

		firstParts := strings.Split(parts[0], ":")
		if len(firstParts) != 2 {
			return nil, fmt.Errorf("failed to parse connection string")
		}
		username := firstParts[0]
		pass := firstParts[1]

		c.addr = parts[1]

		if err := c.establishConnection(); err != nil {
			return nil, err
		}
		_, err := c.Auth(username, pass)
		if err != nil {
			return nil, fmt.Errorf("failed to authenticate")
		}

		return &c, nil
	} else if opt.Addr == "" {
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
