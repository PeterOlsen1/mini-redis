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

func setupAndFlush(t *testing.T) *client.RedisClient {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	_, err = c.FlushAll()
	if err != nil {
		t.Errorf("initial flush")
	}
	return c
}

func teardown(c *client.RedisClient, t *testing.T) {
	err := c.Close()
	if err != nil {
		t.Errorf("client close")
	}
}

func checkExpect[T comparable](resp T, err error, cmd types.Command, expect T, t *testing.T) {
	if err != nil {
		t.Errorf("%s command (%s)", cmd.String(), err)
	}
	if resp != expect {
		t.Errorf("%s not met with %v (%v)", cmd.String(), expect, resp)
	}
}

func checkExpectArray[T comparable](resp []T, err error, cmd types.Command, expect []T, t *testing.T) {
	if err != nil {
		t.Errorf("%s command (%s)", cmd.String(), err)
	}

	if len(resp) != len(expect) {
		t.Errorf("%s slice lengths do not match! %d, %d", cmd.String(), len(resp), len(expect))
	}

	for i := range len(resp) {
		if resp[i] != expect[i] {
			t.Errorf("%s index %d not met with %v (%v)", cmd.String(), i, expect[i], resp[i])
		}
	}
}

func checkError(resp any, err error, cmd types.Command, errText string, t *testing.T) {
	if err == nil {
		t.Errorf("%s command expected error (%s)", cmd.String(), resp)
	}

	if errText != "" && err.Error() != errText {
		t.Errorf("%s error not met with %s (%s)", cmd.String(), errText, err.Error())
	}
}
