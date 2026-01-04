package miniredis_test

import (
	"mini-redis/types"
	"mini-redis/types/errors"
	"testing"
)

func TestLPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	i, err := c.LPush("test", "test")
	checkError(i, err, types.LPUSH, errors.WRONGTYPE, t)

	i, err = c.LPush("test2", "test")
	checkExpect(i, err, types.LPUSH, 1, t)

	i, err = c.LPush("test2", "test")
	checkExpect(i, err, types.LPUSH, 2, t)

	s, err = c.Get("test2")
	checkError(s, err, types.GET, errors.WRONGTYPE, t)
}

func TestRPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	i, err := c.RPush("test", "test")
	checkError(i, err, types.RPUSH, errors.WRONGTYPE, t)

	i, err = c.RPush("test2", "test")
	checkExpect(i, err, types.RPUSH, 1, t)

	i, err = c.RPush("test2", "test")
	checkExpect(i, err, types.RPUSH, 2, t)

	s, err = c.Get("test2")
	checkError(s, err, types.GET, errors.WRONGTYPE, t)
}
