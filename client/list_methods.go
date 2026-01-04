package client

func (c *RedisClient) LPush(key string, values ...string) (int, error) {
	req := InitRequest(2+len(values), "LPUSH").AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.sendAndReceiveInt(
		req,
	)
}

func (c *RedisClient) RPush(key string, values ...string) (int, error) {
	req := InitRequest(2+len(values), "RPUSH").AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.sendAndReceiveInt(
		req,
	)
}
