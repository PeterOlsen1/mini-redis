package client

import (
	"fmt"
)

func (c *RedisClient) sendAndReceive(data string, bufLen int) (string, error) {
	_, err := c.conn.Write([]byte(data))
	if err != nil {
		return "", err
	}

	buf := make([]byte, bufLen)
	n, err := c.conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func (c *RedisClient) Ping() (string, error) {
	return c.sendAndReceive("*1\r\n$4\r\nPING\r\n", 128)
}

func (c *RedisClient) Echo(msg string) (string, error) {
	return c.sendAndReceive(fmt.Sprintf("*1\r\n$%d\r\n%s\r\n", len(msg), msg), len(msg)+24)
}
