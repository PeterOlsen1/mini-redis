package miniredis_test

import (
	"fmt"
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

	i, err = c.LPush("test2", "test")
	checkExpect(i, err, types.LPUSH, 3, t)

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

func TestLPopOne(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	respArr, err := c.LPop("test")
	checkError(respArr, err, types.LPOP, errors.WRONGTYPE, t)

	for i := range 5 {
		resp, err := c.LPush("test2", fmt.Sprintf("test-%d", i))
		checkExpect(resp, err, types.LPUSH, i+1, t)
	}

	expect := []string{"test-4"}
	respArr, err = c.LPop("test2")
	checkExpectArray(respArr, err, types.LPOP, expect, t)
}

func TestLPopMany(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	for i := range 5 {
		resp, err := c.LPush("test2", fmt.Sprintf("test-%d", i))
		checkExpect(resp, err, types.LPUSH, i+1, t)
	}

	expect := []string{"test-4", "test-3", "test-2"}
	sarr, err := c.LPopMany("test2", 3)
	checkExpectArray(sarr, err, types.LPOP, expect, t)
}
