package miniredis_test

import (
	"mini-redis/client"
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
	if err != nil {
		t.Errorf("ping command (%s)", err)
	}
	if s != "PONG" {
		t.Errorf("ping not met with PONG (%s)", s)
	}

	s, err = c.Ping("HELLO")
	if err != nil {
		t.Errorf("ping command (%s)", err)
	}
	if s != "HELLO" {
		t.Errorf("ping not met with HELLO (%s)", s)
	}
}

func TestEcho(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Echo("HELLO")
	if err != nil {
		t.Errorf("ping command (%s)", err)
	}
	if s != "HELLO" {
		t.Errorf("ping not met with HELLO (%s)", s)
	}
}
func TestSet(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Set("test", "TEST")
	if err != nil {
		t.Errorf("set command (%s)", err)
	}
	if s != "OK" {
		t.Errorf("set not met with OK (%s)", s)
	}
}

func TestGet(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Get("test")
	if err != nil {
		t.Errorf("get command (%s)", err)
	}
	if s != "TEST" {
		t.Errorf("get not met with TEST (%s)", s)
	}
}

func TestFlushAll(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.FlushAll()
	if err != nil {
		t.Errorf("flushall command (%s)", err)
	}
	if s != "OK" {
		t.Errorf("get not met with OK (%s)", s)
	}
}

func TestEmptyGet(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()

	s, err := c.Get("test")
	if err != nil {
		t.Errorf("get command (%s)", err)
	}
	if s != "" {
		t.Errorf("get not met with \"\" (%s)", s)
	}
}

func TestExists(t *testing.T) {
	c := setup(t)
	defer teardown(c, t)

	s, err := c.Set("test", "TEST")
	if err != nil {
		t.Errorf("set command (%s)", err)
	}
	if s != "OK" {
		t.Errorf("set not met with OK (%s)", s)
	}

	s, err = c.Set("test2", "TEST")
	if err != nil {
		t.Errorf("set command (%s)", err)
	}
	if s != "OK" {
		t.Errorf("set not met with OK (%s)", s)
	}

	keys := make([]string, 0)
	keys = append(keys, "test")
	s, err = c.Exists(keys)
	if err != nil {
		t.Errorf("set command (%s)", err)
	}
	if s != "1" {
		t.Errorf("set not met with 1 (%s)", s)
	}

	keys = append(keys, "test2")
	s, err = c.Exists(keys)
	if err != nil {
		t.Errorf("set command (%s)", err)
	}
	if s != "2" {
		t.Errorf("set not met with 2 (%s)", s)
	}

}
