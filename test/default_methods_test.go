package miniredis_test

import (
	"mini-redis/client"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	c := setup(t)
	teardown(c, t)
}

func TestPing(t *testing.T) {
	c, err := client.NewClient(&client.ClientOptions{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Ping("")
	checkExpect(s, err, commands.PING, "PONG", t)

	s, err = c.Ping("HELLO")
	checkExpect(s, err, commands.PING, "HELLO", t)
}

func TestEcho(t *testing.T) {
	c, err := client.NewClient(&client.ClientOptions{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Echo("HELLO")
	checkExpect(s, err, commands.ECHO, "HELLO", t)
}

func TestSet(t *testing.T) {
	c, err := client.NewClient(&client.ClientOptions{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)
}

func TestGet(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	s, err = c.Get("test")
	checkExpect(s, err, commands.GET, "TEST", t)
}

func TestDelete(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	s, err = c.Get("test")
	checkExpect(s, err, commands.GET, "TEST", t)

	s, err = c.Del("test")
	checkExpect(s, err, commands.GET, "OK", t)

	s, err = c.Get("test")
	checkExpect(s, err, commands.GET, "", t)
}

func TestFlushAll(t *testing.T) {
	c, err := client.NewClient(&client.ClientOptions{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.FlushAll()
	checkExpect(s, err, commands.FLUSHALL, "OK", t)
}

func TestEmptyGet(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Get("test")
	checkExpect(s, err, commands.GET, "", t)
}

func TestExists(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	s, err = c.Set("test2", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	s, err = c.Exists("test")
	checkExpect(s, err, commands.EXISTS, "1", t)

	s, err = c.Exists("test", "test2")
	checkExpect(s, err, commands.EXISTS, "2", t)
}

func TestExpire(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.Expire("awidawbnd", 10)
	checkExpect(s, err, commands.EXPIRE, "0", t)

	s, err = c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	s, err = c.Expire("test", 10)
	checkExpect(s, err, commands.EXPIRE, "1", t)
}

func TestTTL(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	s, err := c.TTL("test")
	checkExpect(s, err, commands.TTL, "-2", t)

	s, err = c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	s, err = c.TTL("test")
	checkExpect(s, err, commands.TTL, "-1", t)

	s, err = c.Expire("test", 2)
	checkExpect(s, err, commands.EXPIRE, "1", t)

	time.Sleep(500 * time.Millisecond)
	s, err = c.TTL("test")
	checkExpect(s, err, commands.TTL, "1", t)

	time.Sleep(2000 * time.Millisecond)
	s, err = c.TTL("test")
	checkExpect(s, err, commands.TTL, "-2", t)
}

func TestIncr(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	i, err := c.Incr("test")
	checkExpect(i, err, commands.INCR, 1, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	i, err = c.Incr("test")
	checkError(i, err, commands.INCR, errors.NOT_INTEGER.Error(), t)

	s, err = c.Set("test", "1")
	checkExpect(s, err, commands.SET, "OK", t)

	i, err = c.Incr("test")
	checkExpect(i, err, commands.INCR, 2, t)

	s, err = c.Get("test")
	checkExpect(s, err, commands.GET, "2", t)
}

func TestDecr(t *testing.T) {
	c := setupAndFlush(t)
	defer teardown(c, t)

	i, err := c.Decr("test")
	checkExpect(i, err, commands.DECR, -1, t)

	s, err := c.Set("test", "TEST")
	checkExpect(s, err, commands.SET, "OK", t)

	i, err = c.Decr("test")
	checkError(i, err, commands.DECR, errors.NOT_INTEGER.Error(), t)

	s, err = c.Set("test", "1")
	checkExpect(s, err, commands.SET, "OK", t)

	i, err = c.Decr("test")
	checkExpect(i, err, commands.DECR, 0, t)

	s, err = c.Get("test")
	checkExpect(s, err, commands.GET, "0", t)
}
