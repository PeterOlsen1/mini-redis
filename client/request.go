package client

import (
	"fmt"
	"mini-redis/resp"
	"strconv"
)

const RESP_BUF_LEN = 256

func (c *RedisClient) makeRequest(req *RequestBuilder) (any, resp.RespType, error) {
	_, err := c.conn.Write(req.ToBytes())
	if err != nil {
		return "", resp.NULL, err
	}

	buf := make([]byte, RESP_BUF_LEN)
	n, err := c.conn.Read(buf)
	if err != nil {
		return "", resp.NULL, err
	}
	buf = buf[:n]
	result, resType, err := resp.Decode(buf)

	if err != nil {
		return "", resp.NULL, err
	}

	return result, resType, err
}

func (c *RedisClient) sendAndReceive(req *RequestBuilder) (string, error) {
	result, resType, err := c.makeRequest(req)

	if err != nil {
		return "", err
	}

	switch resType {
	case resp.STRING, resp.BULK_STRING, resp.NULL:
		out, ok := result.(string)
		if !ok {
			return "", fmt.Errorf("failed to parse return string")
		}

		if resType == resp.ERR {
			return "", fmt.Errorf("%s", out)
		}

		return out, nil
	case resp.ERR:
		out, ok := result.(string)
		if !ok {
			return "", fmt.Errorf("failed to parse return string")
		}
		return "", fmt.Errorf("%s", out)
	case resp.ARRAY:
		return "", fmt.Errorf("methdo returned array")
	}

	return "", fmt.Errorf("unknown RESP type returned")
}

func (c *RedisClient) sendAndReceiveList(req *RequestBuilder) ([]string, error) {
	result, resType, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	switch resType {
	case resp.STRING, resp.BULK_STRING, resp.NULL, resp.ERR:
		out, ok := result.(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert return string")
		}

		if resType == resp.ERR {
			return nil, fmt.Errorf("%s", out)
		}

		return []string{out}, nil
	case resp.ARRAY:
		out, ok := result.([]string)
		if !ok {
			return nil, fmt.Errorf("failed to convert return string array")
		}
		return out, nil
	}

	return nil, fmt.Errorf("unknown RESP type returned")
}

func (c *RedisClient) sendAndReceiveInt(req *RequestBuilder) (int, error) {
	result, resType, err := c.makeRequest(req)

	if err != nil {
		return 0, err
	}

	switch resType {
	case resp.STRING, resp.BULK_STRING, resp.NULL, resp.ARRAY:
		out, err := strconv.Atoi(result.(string))
		if err != nil {
			return 0, fmt.Errorf("failed to convert return int")
		}

		return out, nil
	case resp.ERR:
		out, ok := result.(string)
		if !ok {
			return 0, fmt.Errorf("failed to parse return string")
		}
		return 0, fmt.Errorf("%s", out)
	}

	return 0, fmt.Errorf("unknown RESP type returned")
}
