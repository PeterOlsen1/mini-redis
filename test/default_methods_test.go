package miniredis_test

import (
	"fmt"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	setup(t)
}

func TestPing(t *testing.T) {
	c := setupAndFlush(t)

	s, err := c.Ping("")
	checkExpect(s, err, commands.PING, "PONG", t)

	s, err = c.Ping("HELLO")
	checkExpect(s, err, commands.PING, "HELLO", t)
}

func TestEcho(t *testing.T) {
	c := setupAndFlush(t)

	s, err := c.Echo("HELLO")
	checkExpect(s, err, commands.ECHO, "HELLO", t)
}

func TestSet(t *testing.T) {
	c := setupAndFlush(t)

	err := c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}
}

func TestGet(t *testing.T) {
	c := setupAndFlush(t)

	err := c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	s, err := c.Get("test")
	checkExpect(s, err, commands.GET, "TEST", t)
}

func TestDelete(t *testing.T) {
	c := setupAndFlush(t)

	err := c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	s, err := c.Get("test")
	checkExpect(s, err, commands.GET, "TEST", t)

	err = c.Del("test")
	if err != nil {
		t.Errorf("DEL command returned an error: %e", err)
	}

	s, err = c.Get("test")
	checkExpect(s, err, commands.GET, "", t)
}

func TestFlushAll(t *testing.T) {
	c := setup(t)

	err := c.FlushAll()
	if err != nil {
		t.Errorf("FLUSHALL command returned an error: %e", err)
	}
}

func TestEmptyGet(t *testing.T) {
	c := setupAndFlush(t)

	s, err := c.Get("test")
	checkExpect(s, err, commands.GET, "", t)
}

func TestExists(t *testing.T) {
	c := setupAndFlush(t)

	err := c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	err = c.Set("test2", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	s, err := c.Exists("test")
	checkExpect(s, err, commands.EXISTS, "1", t)

	s, err = c.Exists("test", "test2")
	checkExpect(s, err, commands.EXISTS, "2", t)
}

func TestExpire(t *testing.T) {
	c := setupAndFlush(t)

	s, err := c.Expire("awidawbnd", 10)
	checkExpect(s, err, commands.EXPIRE, "0", t)

	err = c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	s, err = c.Expire("test", 10)
	checkExpect(s, err, commands.EXPIRE, "1", t)
}

func TestTTL(t *testing.T) {
	c := setupAndFlush(t)

	s, err := c.TTL("test")
	checkExpect(s, err, commands.TTL, "-2", t)

	err = c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

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

	i, err := c.Incr("test")
	checkExpect(i, err, commands.INCR, 1, t)

	err = c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	i, err = c.Incr("test")
	checkError(i, err, commands.INCR, errors.NOT_INTEGER.Error(), t)

	err = c.Set("test", "1")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	i, err = c.Incr("test")
	checkExpect(i, err, commands.INCR, 2, t)

	s, err := c.Get("test")
	checkExpect(s, err, commands.GET, "2", t)
}

func TestDecr(t *testing.T) {
	c := setupAndFlush(t)

	i, err := c.Decr("test")
	checkExpect(i, err, commands.DECR, -1, t)

	err = c.Set("test", "TEST")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	i, err = c.Decr("test")
	checkError(i, err, commands.DECR, errors.NOT_INTEGER.Error(), t)

	err = c.Set("test", "1")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	i, err = c.Decr("test")
	checkExpect(i, err, commands.DECR, 0, t)

	s, err := c.Get("test")
	checkExpect(s, err, commands.GET, "0", t)
}

func TestKeys(t *testing.T) {
	c := setupAndFlush(t)

	setKeys := make([]string, 5)
	for i := range 5 {
		key := fmt.Sprintf("test-%d", i)
		setKeys[i] = key
		err := c.Set(key, "1")
		if err != nil {
			t.Errorf("SET command returned error: %e", err)
		}
	}

	keys, err := c.Keys()
	if err != nil {
		t.Errorf("failed to complete KEYS command: %e", err)
	}
	// cannot check with checkExpectArray method as map key ordering is random
	for i := range 5 {
		found := false
		for j := range 5 {
			if setKeys[j] == keys[i] {
				found = true
			}
		}

		if !found {
			t.Errorf("Key %s not found in KEYS response array", keys[i])
		}
	}
}

func TestExpireAtTime(t *testing.T) {
	c := setupAndFlush(t)

	respInt, err := c.ExpireTime("hello")
	checkExpect(respInt, err, commands.EXPIRETIME, -2, t)

	err = c.Set("hello", "hello")
	if err != nil {
		t.Errorf("SET command returned error: %e", err)
	}

	respInt, err = c.ExpireTime("hello")
	checkExpect(respInt, err, commands.EXPIRETIME, -1, t)

	newTime := time.Now().UnixMilli()/1000 + 100
	respInt, err = c.ExpireAt("hello", int(newTime))
	checkExpect(respInt, err, commands.EXPIREAT, 1, t)

	respInt, err = c.ExpireTime("hello")
	checkExpect(respInt, err, commands.EXPIRETIME, int(newTime), t)
}
