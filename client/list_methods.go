package client

func (c *RedisClient) LPush(key string, values ...string) (string, error) {
	req := InitRequest(2+len(values), "LPUSH").AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.sendAndReceive(
		req,
	)
}
