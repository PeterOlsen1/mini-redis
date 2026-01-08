package client

import "mini-redis/types/commands"

// Ping function. Pass in an empty string for no message
func (c *RedisClient) Ping(msg string) (string, error) {
	if msg == "" {
		return c.SendAndReceive(InitRequest(commands.PING))
	}

	return c.SendAndReceive(
		InitRequest(commands.PING).AddParam(msg),
	)
}

func (c *RedisClient) Echo(msg string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.ECHO).AddParam(msg),
	)
}

func (c *RedisClient) Set(key string, value string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.SET).AddParam(key).AddParam(value),
	)
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.GET).AddParam(key),
	)
}

func (c *RedisClient) Del(key string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.DEL).AddParam(key),
	)
}

func (c *RedisClient) Exists(keys ...string) (string, error) {
	req := InitRequest(commands.EXISTS)
	for _, key := range keys {
		req.AddParam(key)
	}

	return c.SendAndReceive(
		req,
	)
}

func (c *RedisClient) Expire(key string, seconds int) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.EXPIRE).AddParam(key).AddParamInt(seconds),
	)
}

func (c *RedisClient) TTL(key string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.TTL).AddParam(key),
	)
}

func (c *RedisClient) Incr(key string) (int, error) {
	return c.SendAndReceiveInt(
		InitRequest(commands.INCR).AddParam(key),
	)
}

func (c *RedisClient) Decr(key string) (int, error) {
	return c.SendAndReceiveInt(
		InitRequest(commands.DECR).AddParam(key),
	)
}

func (c *RedisClient) FlushAll() (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.FLUSHALL),
	)
}

func (c *RedisClient) Info() (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.INFO).SetBufSize(4096),
	)
}

func (c *RedisClient) Keys() ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest(commands.KEYS).SetBufSize(4096),
	)
}

func (c *RedisClient) ExpireAt(key string, secs int) (int, error) {
	return c.SendAndReceiveInt(
		InitRequest(commands.EXPIREAT).AddParam(key).AddParamInt(secs),
	)
}

func (c *RedisClient) ExpireTime(key string) (int, error) {
	return c.SendAndReceiveInt(
		InitRequest(commands.EXPIRETIME).AddParam(key),
	)
}
