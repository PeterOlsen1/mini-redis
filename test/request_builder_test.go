package miniredis_test

import (
	"bytes"
	"mini-redis/client"
	"testing"
)

func TestRequestBuilder(t *testing.T) {
	req := client.InitRequest("PING")

	expect := "*1\r\n$4\r\nPING\r\n"
	if req.String() != expect {
		t.Errorf("PING req does not match expected\nGot: %s\nExpected: %s\n", req.String(), expect)
	}

	if !bytes.Equal(req.ToBytes(), []byte(expect)) {
		t.Errorf("PING bytes are not equal")
	}

	req = req.AddParam("TEST")
	expect = "*2\r\n$4\r\nPING\r\n$4\r\nTEST\r\n"
	if req.String() != expect {
		t.Errorf("PING req does not match expected\nGot: %s\nExpected: %s\n", req.String(), expect)
	}

	if !bytes.Equal(req.ToBytes(), []byte(expect)) {
		t.Errorf("PING bytes are not equal")
	}

	req = req.AddParamInt(1)
	expect = "*3\r\n$4\r\nPING\r\n$4\r\nTEST\r\n$1\r\n1\r\n"
	if req.String() != expect {
		t.Errorf("PING req does not match expected\nGot: %s\nExpected: %s\n", req.String(), expect)
	}

	if !bytes.Equal(req.ToBytes(), []byte(expect)) {
		t.Errorf("PING bytes are not equal")
	}
}
