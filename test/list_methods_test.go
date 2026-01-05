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

	respInt, err := c.LPush("test", "test")
	checkError(respInt, err, types.LPUSH, errors.WRONGTYPE, t)

	for i := range 3 {
		respInt, err = c.LPush("test2", "test")
		checkExpect(respInt, err, types.LPUSH, i+1, t)
	}

	s, err = c.Get("test2")
	checkError(s, err, types.GET, errors.WRONGTYPE, t)
}

func TestRPush(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	respInt, err := c.RPush("test", "test")
	checkError(respInt, err, types.RPUSH, errors.WRONGTYPE, t)

	for i := range 3 {
		respInt, err = c.RPush("test2", "test")
		checkExpect(respInt, err, types.RPUSH, i+1, t)
	}

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
		respInt, err := c.LPush("test2", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, types.LPUSH, i+1, t)
	}

	expect := []string{"test-4"}
	respArr, err = c.LPop("test2")
	checkExpectArray(respArr, err, types.LPOP, expect, t)
}

func TestLPopMany(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	for i := range 5 {
		respInt, err := c.LPush("test2", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, types.LPUSH, i+1, t)
	}

	expect := []string{"test-4", "test-3", "test-2"}
	respArr, err := c.LPopMany("test2", 3)
	fmt.Printf("RESP array: %v\n", respArr)
	checkExpectArray(respArr, err, types.LPOP, expect, t)
}

func TestLRange(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	for i := range 5 {
		respInt, err := c.LPush("test", fmt.Sprintf("test-%d", i))
		checkExpect(respInt, err, types.LPUSH, i+1, t)
	}

	// expect := []string{"test-1", "test-2"}
	// arr, err := c.
	// checkExpectArray()
}
