package miniredis_test

import (
	"mini-redis/client"
	"mini-redis/types/commands"
	"testing"
)

var globalClient *client.RedisClient

func setup(t *testing.T) *client.RedisClient {
	if globalClient == nil {
		c, err := client.NewClient(&client.ClientOptions{
			URL: "redis://admin:admin@localhost:6379",
		})
		if err != nil {
			t.Errorf("client initialization")
		}

		globalClient = c
		return c
	}

	return globalClient
}

func setupAndFlush(t *testing.T) *client.RedisClient {
	if globalClient == nil {
		c, err := client.NewClient(&client.ClientOptions{
			URL: "redis://admin:admin@localhost:6379",
		})
		if err != nil {
			t.Errorf("client initialization")
		}

		_, err = c.FlushAll()
		if err != nil {
			t.Errorf("initial flush")
		}

		globalClient = c
		return c
	}

	_, err := globalClient.FlushAll()
	if err != nil {
		t.Errorf("initial flush")
	}
	return globalClient
}

// teardown method left empty for now due to use of global client
func teardown() {
	if globalClient != nil {
		globalClient.Close()
	}
}

func checkExpect[T comparable](resp T, err error, cmd commands.Command, expect T, t *testing.T) {
	if err != nil {
		t.Errorf("%s command (%s)", cmd.String(), err)
	}
	if resp != expect {
		t.Errorf("%s not met with %v (%v)", cmd.String(), expect, resp)
	}
}

func checkExpectArray[T comparable](resp []T, err error, cmd commands.Command, expect []T, t *testing.T) {
	if err != nil {
		t.Errorf("%s command (%s)", cmd.String(), err)
	}

	if len(resp) != len(expect) {
		// fatal becuase next lines will cause panic
		t.Fatalf("%s slice lengths do not match! expected: %d, got: %d", cmd.String(), len(expect), len(resp))
	}

	for i := range len(resp) {
		if resp[i] != expect[i] {
			t.Errorf("%s index %d not met with %v (%v)", cmd.String(), i, expect[i], resp[i])
		}
	}
}

func checkError(resp any, err error, cmd commands.Command, errText string, t *testing.T) {
	if err == nil {
		t.Errorf("%s command expected error (%s)", cmd.String(), resp)
	}

	if errText != "" && err.Error() != errText {
		t.Errorf("%s error not met with %s (%s)", cmd.String(), errText, err.Error())
	}
}
