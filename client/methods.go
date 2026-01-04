package client

import (
	"strconv"
)

// Ping function. Pass in an empty string for no message
func (c *RedisClient) Ping(msg string) (string, error) {
	if msg == "" {
		return c.sendAndReceive(InitRequest(1, "PING"))
	}

	return c.sendAndReceive(
		InitRequest(2, "PING").AddParam(msg),
	)
}

func (c *RedisClient) Echo(msg string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "ECHO").AddParam(msg),
	)
}

func (c *RedisClient) Set(key string, value string) (string, error) {
	return c.sendAndReceive(
		InitRequest(3, "SET").AddParam(key).AddParam(value),
	)
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "GET").AddParam(key),
	)
}

func (c *RedisClient) Del(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "DEL").AddParam(key),
	)
}

func (c *RedisClient) Exists(keys ...string) (string, error) {
	req := InitRequest(1+len(keys), "EXISTS")
	for _, key := range keys {
		req.AddParam(key)
	}

	return c.sendAndReceive(
		req,
	)
}

func (c *RedisClient) Expire(key string, seconds int) (string, error) {
	return c.sendAndReceive(
		InitRequest(3, "EXPIRE").AddParam(key).AddParam(strconv.Itoa(seconds)),
	)
}

func (c *RedisClient) TTL(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest(2, "TTL").AddParam(key),
	)
}

func (c *RedisClient) Incr(key string) (int, error) {
	return c.sendAndReceiveInt(
		InitRequest(2, "INCR").AddParam(key),
	)
}

func (c *RedisClient) Decr(key string) (int, error) {
	return c.sendAndReceiveInt(
		InitRequest(2, "DECR").AddParam(key),
	)
}

func (c *RedisClient) FlushAll() (string, error) {
	return c.sendAndReceive(
		InitRequest(1, "FLUSHALL"),
	)
}
