package jsonutil

import "github.com/buger/jsonparser"

func JSONGetString(data []byte, keys ...string) (string, error) {
	return jsonparser.GetString(data, keys...)
}

func JSONMustGetString(data []byte, keys ...string) string {
	s, err := JSONGetString(data, keys...)
	if err != nil {
		return ""
	}
	return s
}

func JSONGetFloat(data []byte, keys ...string) (float64, error) {
	return jsonparser.GetFloat(data, keys...)
}

func JSONMustGetFloat(data []byte, keys ...string) float64 {
	f, err := JSONGetFloat(data, keys...)
	if err != nil {
		return 0.0
	}
	return f
}

func JSONGetInt(data []byte, keys ...string) (int64, error) {
	return jsonparser.GetInt(data, keys...)
}

func JSONMustGetInt(data []byte, keys ...string) int64 {
	i, err := JSONGetInt(data, keys...)
	if err != nil {
		return 0
	}
	return i
}

func JSONGetBoolean(data []byte, keys ...string) (bool, error) {
	return jsonparser.GetBoolean(data, keys...)
}

func JSONMustGetBoolean(data []byte, keys ...string) bool {
	b, err := JSONGetBoolean(data, keys...)
	if err != nil {
		return false
	}
	return b
}
