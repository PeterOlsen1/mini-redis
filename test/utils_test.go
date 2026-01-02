package miniredis_test

import (
	"mini-redis/client"
	"mini-redis/types"
	"testing"
)

func setup(t *testing.T) *client.RedisClient {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	return c
}

func teardown(c *client.RedisClient, t *testing.T) {
	err := c.Close()
	if err != nil {
		t.Errorf("client close")
	}
}

func checkExpected(resp string, err error, cmd types.Command, expect string, t *testing.T) {
	if err != nil {
		t.Errorf("%s command (%s)", cmd.String(), err)
	}
	if resp != expect {
		t.Errorf("%s not met with %s (%s)", cmd.String(), expect, resp)
	}
}
