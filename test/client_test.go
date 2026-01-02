package miniredis_test

import (
	"mini-redis/client"
	"testing"
)

func TestInit(t *testing.T) {
	c, err := client.NewClient(&client.Options{Addr: "localhost:6379"})
	if err != nil {
		t.Errorf("client initialization")
	}

	defer c.Close()
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

	s, err := c.Set("test", "test")
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
	if s != "" {
		t.Errorf("get not met with \"\" (%s)", s)
	}

	// set test key
	c.Set("test", "TEST")
	if err != nil {
		t.Errorf("set command (%s)", err)
	}

	s, err = c.Get("test")
	if err != nil {
		t.Errorf("get command (%s)", err)
	}
	if s != "TEST" {
		t.Errorf("get not met with TEST (%s)", s)
	}
}
