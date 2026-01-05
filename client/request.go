package client

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/types/commands"
	"strconv"
	"strings"
)

func InitRequest(command commands.Command) *RequestBuilder {
	cmd := command.String()
	return &RequestBuilder{
		req: fmt.Sprintf("\r\n$%d\r\n%s\r\n", len(cmd), cmd),
		len: 1,
	}
}

func (r *RequestBuilder) AddParam(param string) *RequestBuilder {
	r.req += fmt.Sprintf("$%d\r\n%s\r\n", len(param), param)
	r.len += 1
	return r
}

func (r *RequestBuilder) AddParamInt(param int) *RequestBuilder {
	pString := strconv.Itoa(param)
	r.req += fmt.Sprintf("$%d\r\n%s\r\n", len(pString), pString)
	r.len += 1
	return r
}

func (r *RequestBuilder) ToBytes() []byte {
	return []byte(r.String())
}

func (r *RequestBuilder) String() string {
	return fmt.Sprintf("*%d%s", r.len, r.req)
}

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

// Lower level function that is exposed to other redis client methods.
// This is intended to be used if you want to create arbitrary requests
func (c *RedisClient) SendAndReceive(req *RequestBuilder) (string, error) {
	result, resType, err := c.makeRequest(req)

	if err != nil {
		return "", err
	}

	switch resType {
	case resp.NULL:
		return "", nil
	case resp.STRING, resp.BULK_STRING:
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
		out, ok := result.([]string)
		if !ok {
			return "", fmt.Errorf("failed to parse return string array")
		}

		outStrings := make([]string, len(out))
		for i := range len(out) {
			outStrings[i] = fmt.Sprintf("%d) \"%s\"", i, out[i])
		}
		return strings.Join(outStrings, "\n"), nil
	}

	return "", fmt.Errorf("unknown RESP type returned")
}

func (c *RedisClient) SendAndReceiveList(req *RequestBuilder) ([]string, error) {
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

func (c *RedisClient) SendAndReceiveInt(req *RequestBuilder) (int, error) {
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
