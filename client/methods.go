package client

// Ping function. Pass in an empty string for no message
func (c *RedisClient) Ping(msg string) (string, error) {
	if msg == "" {
		return c.sendAndReceive(InitRequest("PING"))
	}

	return c.sendAndReceive(
		InitRequest("PING").AddParam(msg),
	)
}

func (c *RedisClient) Echo(msg string) (string, error) {
	return c.sendAndReceive(
		InitRequest("ECHO").AddParam(msg),
	)
}

func (c *RedisClient) Set(key string, value string) (string, error) {
	return c.sendAndReceive(
		InitRequest("SET").AddParam(key).AddParam(value),
	)
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest("GET").AddParam(key),
	)
}

func (c *RedisClient) Del(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest("DEL").AddParam(key),
	)
}

func (c *RedisClient) Exists(keys ...string) (string, error) {
	req := InitRequest("EXISTS")
	for _, key := range keys {
		req.AddParam(key)
	}

	return c.sendAndReceive(
		req,
	)
}

func (c *RedisClient) Expire(key string, seconds int) (string, error) {
	return c.sendAndReceive(
		InitRequest("EXPIRE").AddParam(key).AddParamInt(seconds),
	)
}

func (c *RedisClient) TTL(key string) (string, error) {
	return c.sendAndReceive(
		InitRequest("TTL").AddParam(key),
	)
}

func (c *RedisClient) Incr(key string) (int, error) {
	return c.sendAndReceiveInt(
		InitRequest("INCR").AddParam(key),
	)
}

func (c *RedisClient) Decr(key string) (int, error) {
	return c.sendAndReceiveInt(
		InitRequest("DECR").AddParam(key),
	)
}

func (c *RedisClient) FlushAll() (string, error) {
	return c.sendAndReceive(
		InitRequest("FLUSHALL"),
	)
}
