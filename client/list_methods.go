package client

import "strconv"

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

func (c *RedisClient) LPop(key string) ([]string, error) {
	return c.sendAndReceiveList(
		InitRequest(2, "LPOP").AddParam(key),
	)
}

func (c *RedisClient) LPopMany(key string, num int) ([]string, error) {
	return c.sendAndReceiveList(
		InitRequest(3, "LPOP").AddParam(key).AddParam(strconv.Itoa(num)),
	)
}
