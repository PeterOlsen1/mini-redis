package miniredis_test

import (
	"mini-redis/types"
	"testing"
)

func TestLPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.LPush("test", "test")
	checkError(s, err, types.LPUSH, "Operation against a key holding the wrong kind of value", t)

	s, err = c.LPush("test2", "test")
	checkExpect(s, err, types.LPUSH, "1", t)

	s, err = c.LPush("test2", "test")
	checkExpect(s, err, types.LPUSH, "2", t)

	s, err = c.Get("test2")
	checkError(s, err, types.GET, "Operation against a key holding the wrong kind of value", t)
}

func TestRPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.RPush("test", "test")
	checkError(s, err, types.RPUSH, "Operation against a key holding the wrong kind of value", t)

	s, err = c.RPush("test2", "test")
	checkExpect(s, err, types.RPUSH, "1", t)

	s, err = c.RPush("test2", "test")
	checkExpect(s, err, types.RPUSH, "2", t)

	s, err = c.Get("test2")
	checkError(s, err, types.GET, "Operation against a key holding the wrong kind of value", t)
}
