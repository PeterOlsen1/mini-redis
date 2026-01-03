package resp

func Serialize(value any) ([]byte, error)

func Decode(data []byte) (any, error)
