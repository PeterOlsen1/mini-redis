package miniredis_test

import (
	"mini-redis/client"
	"mini-redis/types"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	c := setup(t)
	teardown(c, t)
}

func TestPing(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Ping("")
	checkExpect(s, err, types.PING, "PONG", t)

	s, err = c.Ping("HELLO")
	checkExpect(s, err, types.PING, "HELLO", t)
}

func TestEcho(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Echo("HELLO")
	checkExpect(s, err, types.ECHO, "HELLO", t)
}

func TestSet(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)
}

func TestGet(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Get("test")
	checkExpect(s, err, types.GET, "TEST", t)
}

func TestFlushAll(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.FlushAll()
	checkExpect(s, err, types.FLUSHALL, "OK", t)
}

func TestEmptyGet(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Get("test")
	checkExpect(s, err, types.GET, "", t)
}

func TestExists(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Set("test2", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Exists("test")
	checkExpect(s, err, types.EXISTS, "1", t)

	s, err = c.Exists("test", "test2")
	checkExpect(s, err, types.EXISTS, "2", t)
}

func TestExpire(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Expire("awidawbnd", 10)
	checkExpect(s, err, types.EXPIRE, "0", t)

	s, err = c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Expire("test", 10)
	checkExpect(s, err, types.EXPIRE, "1", t)
}

func TestTTL(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.TTL("test")
	checkExpect(s, err, types.TTL, "-2", t)

	s, err = c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.TTL("test")
	checkExpect(s, err, types.TTL, "-1", t)

	s, err = c.Expire("test", 2)
	checkExpect(s, err, types.EXPIRE, "1", t)

	time.Sleep(500 * time.Millisecond)
	s, err = c.TTL("test")
	checkExpect(s, err, types.TTL, "1", t)

	time.Sleep(2000 * time.Millisecond)
	s, err = c.TTL("test")
	checkExpect(s, err, types.TTL, "-2", t)
}

func TestIncr(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Incr("test")
	checkExpect(s, err, types.INCR, "1", t)

	s, err = c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Incr("test")
	checkError(s, err, types.INCR, "value is not an integer or out of range", t)

	s, err = c.Set("test", "1")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Incr("test")
	checkExpect(s, err, types.INCR, "2", t)

	s, err = c.Get("test")
	checkExpect(s, err, types.GET, "2", t)
}

func TestDecr(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Decr("test")
	checkExpect(s, err, types.DECR, "-1", t)

	s, err = c.Set("test", "TEST")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Decr("test")
	checkError(s, err, types.DECR, "value is not an integer or out of range", t)

	s, err = c.Set("test", "1")
	checkExpect(s, err, types.SET, "OK", t)

	s, err = c.Decr("test")
	checkExpect(s, err, types.DECR, "0", t)

	s, err = c.Get("test")
	checkExpect(s, err, types.GET, "0", t)
}

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
