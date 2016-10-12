package jsonutil

import (
	"github.com/buger/jsonparser"
	"strconv"
)

func GetString(data []byte, keys ...string) (string, error) {
	return jsonparser.GetString(data, keys...)
}

func MustGetString(data []byte, keys ...string) string {
	s, err := GetString(data, keys...)
	if err != nil {
		return ""
	}
	return s
}

func GetFloat(data []byte, keys ...string) (float64, error) {
	return jsonparser.GetFloat(data, keys...)
}

func MustGetFloat(data []byte, keys ...string) float64 {
	f, err := GetFloat(data, keys...)
	if err != nil {
		return 0.0
	}
	return f
}

func GetInt(data []byte, keys ...string) (int64, error) {
	return jsonparser.GetInt(data, keys...)
}

func MustGetInt(data []byte, keys ...string) int64 {
	i, err := GetInt(data, keys...)
	if err != nil {
		if s := MustGetString(data, keys...); len(s) == 0 {
			return 0
		} else if i, err = strconv.ParseInt(s, 10, 64); err != nil {
			return 0
		}
	}
	return i
}

func GetBoolean(data []byte, keys ...string) (bool, error) {
	return jsonparser.GetBoolean(data, keys...)
}

func MustGetBoolean(data []byte, keys ...string) bool {
	b, err := GetBoolean(data, keys...)
	if err != nil {
		return false
	}
	return b
}
