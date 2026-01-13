package client

import "mini-redis/types/commands"

func (c *RedisClient) Auth(user string, password string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.AUTH).AddParam(user).AddParam(password),
	)
}

func (c *RedisClient) Logout() (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.LOGOUT),
	)
}

func (c *RedisClient) Whoami() (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.WHOAMI),
	)
}

func (c *RedisClient) Rmuser(user string) (string, error) {
	return c.SendAndReceive(
		InitRequest(commands.RMUSER).AddParam(user),
	)
}

func (c *RedisClient) UserGet(users ...string) (string, error) {
	req := InitRequest(commands.UGET)
	for _, user := range users {
		req.AddParam(user)
	}

	return c.SendAndReceive(
		req,
	)
}
