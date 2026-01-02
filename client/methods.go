package client

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (c *RedisClient) sendAndReceive(req *RequestBuilder, bufLen int) (string, error) {
	_, err := c.conn.Write([]byte(req.req))
	if err != nil {
		return "", err
	}

	buf := make([]byte, bufLen)
	n, err := c.conn.Read(buf)
	if err != nil {
		return "", err
	}

	// TODO: only implements basic string / error data type reponses. do we need more?
	ret := string(buf[:n])
	if ret[0] == '+' {
		ret = strings.TrimPrefix(ret, "+")
		ret = strings.TrimSuffix(ret, "\r\n")
		return ret, nil
	}

	ret = strings.TrimPrefix(ret, "-ERR ")
	ret = strings.TrimSuffix(ret, "\r\n")
	return "", fmt.Errorf("%s", ret)
}

// Ping function. Pass in an empty string for no message
func (c *RedisClient) Ping(msg string) (string, error) {
	if msg == "" {
		return c.sendAndReceive(InitRequest(1, "PING"), 128)
	}

	return c.sendAndReceive(
		InitRequest(2, "PING").AddParam(msg),
		len(msg)+32,
	)
}

func (c *RedisClient) Echo(msg string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "ECHO").AddParam(msg),
		len(msg)+32,
	)
}

func (c *RedisClient) Set(key string, value string) (string, error) {
	return c.sendAndReceive(
		InitRequest(3, "SET").AddParam(key).AddParam(value),
		len(key)+len(value)+32,
	)
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "GET").AddParam(key),
		len(key)+32,
	)
}

func (c *RedisClient) Del(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "DEL").AddParam(key),
		len(key)+32,
	)
}

func (c *RedisClient) Exists(keys []string) (string, error) {
	req := InitRequest(1+len(keys), "EXISTS")
	for _, key := range keys {
		req.AddParam(key)
	}

	return c.sendAndReceive(
		req,
		128,
	)
}

func (c *RedisClient) Expire(key string, seconds int) (string, error) {
	return c.sendAndReceive(
		InitRequest(3, "EXPIRE").AddParam(key).AddParam(strconv.Itoa(seconds)),
		128,
	)
}

func (c *RedisClient) TTL(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "TTL").AddParam(key),
		128,
	)
}

func (c *RedisClient) Incr(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "INCR").AddParam(key),
		128,
	)
}

func (c *RedisClient) Decr(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "DECR").AddParam(key),
		128,
	)
}

func (c *RedisClient) FlushAll() (string, error) {
	return c.sendAndReceive(
		InitRequest(1, "FLUSHALL"),
		128,
	)
}
