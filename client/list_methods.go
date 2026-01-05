package client

func (c *RedisClient) LPush(key string, values ...string) (int, error) {
	req := InitRequest("LPUSH").AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.SendAndReceiveInt(
		req,
	)
}

func (c *RedisClient) RPush(key string, values ...string) (int, error) {
	req := InitRequest("RPUSH").AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.SendAndReceiveInt(
		req,
	)
}

func (c *RedisClient) LPop(key string) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest("LPOP").AddParam(key),
	)
}

func (c *RedisClient) LPopMany(key string, num int) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest("LPOP").AddParam(key).AddParamInt(num),
	)
}

func (c *RedisClient) LRange(key string, start int, end int) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest("LRANGE").AddParam(key).AddParamInt(start).AddParamInt(end),
	)
}
