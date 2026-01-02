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

	c.Ping()
}
