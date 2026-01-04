package resp

import (
	"fmt"
	"strconv"
	"strings"
)

// Implementation of the RESP - REdis Serialization Protocol
// Somewhat naive, as bulk strings are not handled properly (length)

func Serialize(value any, valueType RespType) ([]byte, error) {
	switch valueType {
	case STRING:
		return fmt.Appendf(nil, "+%s\r\n", value), nil
	case ERR:
		return fmt.Appendf(nil, "-%s\r\n", value), nil
	case NULL:
		return []byte("_\r\n"), nil
	case ARRAY: // assume array will only be made up of bulk strings
		var out strings.Builder
		out.WriteString("*")
		valueArr, ok := value.([]string)
		if !ok {
			return nil, fmt.Errorf("could not convert value to array")
		}
		out.WriteString(strconv.Itoa(len(valueArr)) + "\r\n")

		for _, s := range valueArr {
			fmt.Fprintf(&out, "$%d\r\n%s\r\n", len(s), s)
		}
		return []byte(out.String()), nil
	case BULK_STRING:
		strVal, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("could not convert value to string")
		}
		return fmt.Appendf(nil, "$%d\r\n%s\r\n", len(strVal), strVal), nil
	}

	return nil, fmt.Errorf("invalid valueType")
}

func Decode(data []byte) (any, RespType, error) {
	if len(data) == 0 {
		return nil, NULL, fmt.Errorf("byte array is empty")
	}

	strData := string(data)
	switch data[0] {
	case '+':
		middleData := strings.TrimSuffix(strData[1:], "\r\n")
		return middleData, STRING, nil
	case '-':
		middleData := strings.TrimSuffix(strData[1:], "\r\n")
		return middleData, ERR, nil
	case '_':
		return nil, NULL, nil
	case '*':
		strList := strings.Split(strData, "\r\n")
		header := strList[0]
		listLenStr := strings.TrimPrefix(strings.TrimSuffix(header, "\r\n"), "*")
		listLen, err := strconv.Atoi(listLenStr)
		if err != nil {
			return nil, ERR, fmt.Errorf("failed to gather list length")
		}

		out := make([]string, listLen)
		for i := range listLen {
			j := (i * 2) + 1

			content := strList[j+1]
			content = strings.TrimSuffix(content, "\r\n")

			out[i] = content
		}
		return out, ARRAY, nil
	case '$':
		strList := strings.Split(strData, "\r\n")

		content := strList[1]
		content = strings.TrimSuffix(content, "\r\n")
		return content, BULK_STRING, nil
	}
	return nil, NULL, fmt.Errorf("could not determine response type")
}
