package client

// Ping function. Pass in an empty string for no message
func (c *RedisClient) Ping(msg string) (string, error) {
	if msg == "" {
		return c.SendAndReceive(InitRequest("PING"))
	}

	return c.SendAndReceive(
		InitRequest("PING").AddParam(msg),
	)
}

func (c *RedisClient) Echo(msg string) (string, error) {
	return c.SendAndReceive(
		InitRequest("ECHO").AddParam(msg),
	)
}

func (c *RedisClient) Set(key string, value string) (string, error) {
	return c.SendAndReceive(
		InitRequest("SET").AddParam(key).AddParam(value),
	)
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.SendAndReceive(
		InitRequest("GET").AddParam(key),
	)
}

func (c *RedisClient) Del(key string) (string, error) {
	return c.SendAndReceive(
		InitRequest("DEL").AddParam(key),
	)
}

func (c *RedisClient) Exists(keys ...string) (string, error) {
	req := InitRequest("EXISTS")
	for _, key := range keys {
		req.AddParam(key)
	}

	return c.SendAndReceive(
		req,
	)
}

func (c *RedisClient) Expire(key string, seconds int) (string, error) {
	return c.SendAndReceive(
		InitRequest("EXPIRE").AddParam(key).AddParamInt(seconds),
	)
}

func (c *RedisClient) TTL(key string) (string, error) {
	return c.SendAndReceive(
		InitRequest("TTL").AddParam(key),
	)
}

func (c *RedisClient) Incr(key string) (int, error) {
	return c.SendAndReceiveInt(
		InitRequest("INCR").AddParam(key),
	)
}

func (c *RedisClient) Decr(key string) (int, error) {
	return c.SendAndReceiveInt(
		InitRequest("DECR").AddParam(key),
	)
}

func (c *RedisClient) FlushAll() (string, error) {
	return c.SendAndReceive(
		InitRequest("FLUSHALL"),
	)
}
