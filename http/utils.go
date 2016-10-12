package http

import (
	"io/ioutil"
	"net/http"
)

func ByteResponseBody(resp *http.Response) ([]byte, error) {
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return p, err
}

func StringResponseBody(resp *http.Response) (string, error) {
	p, err := ByteResponseBody(resp)
	if err != nil {
		return "", err
	}
	return string(p), err
}
