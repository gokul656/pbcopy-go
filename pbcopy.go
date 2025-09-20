package pbcopy

func ReadAll() (string, error) {
	return readAll()
}

func Write(data []byte) (int, error) {
	return write(data)
}

var Unsupported bool
