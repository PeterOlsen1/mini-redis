package client

import "mini-redis/types/commands"

func (c *RedisClient) LPush(key string, values ...string) (int, error) {
	req := InitRequest(commands.LPUSH).AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.SendAndReceiveInt(
		req,
	)
}

func (c *RedisClient) RPush(key string, values ...string) (int, error) {
	req := InitRequest(commands.RPUSH).AddParam(key)
	for _, value := range values {
		req.AddParam(value)
	}

	return c.SendAndReceiveInt(
		req,
	)
}

func (c *RedisClient) LPop(key string) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest(commands.LPOP).AddParam(key),
	)
}

func (c *RedisClient) LPopMany(key string, num int) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest(commands.LPOP).AddParam(key).AddParamInt(num),
	)
}

func (c *RedisClient) LRange(key string, start int, end int) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest(commands.LRANGE).AddParam(key).AddParamInt(start).AddParamInt(end),
	)
}

func (c *RedisClient) LGet(key string) ([]string, error) {
	return c.SendAndReceiveList(
		InitRequest(commands.LGET).AddParam(key),
	)
}
