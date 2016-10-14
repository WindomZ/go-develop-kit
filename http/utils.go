package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ByteRequestBody(req *http.Request) (data []byte, err error) {
	data, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	defer req.Body.Close()
	return
}

func StringRequestBody(req *http.Request) (string, error) {
	data, err := ByteRequestBody(req)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func ByteRequestBodyNoClose(req *http.Request) (data []byte, err error) {
	data, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return
}

func StringRequestBodyNoClose(req *http.Request) (string, error) {
	data, err := ByteRequestBodyNoClose(req)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func ByteResponseBody(resp *http.Response) (data []byte, err error) {
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func StringResponseBody(resp *http.Response) (string, error) {
	data, err := ByteResponseBody(resp)
	if err != nil {
		return "", err
	}
	return string(data), err
}
