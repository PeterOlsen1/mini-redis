package miniredis_test

import (
	"fmt"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"testing"
)

func TestLPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	respInt, err := c.LPush("test", "test")
	checkError(respInt, err, commands.LPUSH, errors.WRONGTYPE.Error(), t)

	for i := range 3 {
		respInt, err = c.LPush("test2", "test")
		checkExpect(respInt, err, commands.LPUSH, i+1, t)
	}

	s, err = c.Get("test2")
	checkError(s, err, commands.GET, errors.WRONGTYPE.Error(), t)
}

func TestRPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	respInt, err := c.RPush("test", "test")
	checkError(respInt, err, commands.RPUSH, errors.WRONGTYPE.Error(), t)

	for i := range 3 {
		respInt, err = c.RPush("test2", "test")
		checkExpect(respInt, err, commands.RPUSH, i+1, t)
	}

	s, err = c.Get("test2")
	checkError(s, err, commands.GET, errors.WRONGTYPE.Error(), t)
}

func TestLPopOne(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	respArr, err := c.LPop("test")
	checkError(respArr, err, commands.LPOP, errors.WRONGTYPE.Error(), t)

	for i := range 5 {
		respInt, err := c.LPush("test2", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, commands.LPUSH, i+1, t)
	}

	expect := []string{"test-4"}
	respArr, err = c.LPop("test2")
	checkExpectArray(respArr, err, commands.LPOP, expect, t)
}

func TestLPopMany(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	for i := range 5 {
		respInt, err := c.LPush("test2", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, commands.LPUSH, i+1, t)
	}

	expect := []string{"test-4", "test-3", "test-2"}
	respArr, err := c.LPopMany("test2", 3)
	fmt.Printf("RESP array: %v\n", respArr)
	checkExpectArray(respArr, err, commands.LPOP, expect, t)
}

func TestLRange(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	for i := range 5 {
		respInt, err := c.LPush("test", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, commands.LPUSH, i+1, t)
	}

	expect := []string{"test-4", "test-3"}
	arr, err := c.LRange("test", 0, 1)
	checkExpectArray(arr, err, commands.LRANGE, expect, t)

	expect = []string{"test-4", "test-3", "test-2", "test-1", "test-0"}
	arr, err = c.LRange("test", -100, 1000)
	checkExpectArray(arr, err, commands.LRANGE, expect, t)

	expect = []string{}
	arr, err = c.LRange("test", 3, 2)
	checkExpectArray(arr, err, commands.LRANGE, expect, t)

	expect = []string{"test-2"}
	arr, err = c.LRange("test", 2, 2)
	checkExpectArray(arr, err, commands.LRANGE, expect, t)
}

func TestLGet(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	for i := range 5 {
		respInt, err := c.LPush("test", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, commands.LPUSH, i+1, t)
	}

	expect := []string{"test-4", "test-3", "test-2", "test-1", "test-0"}
	arr, err := c.LGet("test")
	checkExpectArray(arr, err, commands.LRANGE, expect, t)
}
