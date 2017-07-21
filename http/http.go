package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type HttpClient struct {
	client *http.Client
	Header map[string]string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{client: &http.Client{}}
}

func NewHttpClientWithCAs(cas ...string) *HttpClient {
	pool := x509.NewCertPool()
	for _, ca := range cas {
		if _ca, err := ioutil.ReadFile(ca); err == nil {
			pool.AppendCertsFromPEM(_ca)
		}
	}
	if len(pool.Subjects()) != 0 {
		return &HttpClient{client: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		}}}
	}
	return NewHttpClient()
}

func (s *HttpClient) Clear() {
	s.Header = nil
}

func (s *HttpClient) SetHeader(key, value string) *HttpClient {
	if len(key) == 0 {
		return s
	} else if s.Header == nil {
		s.Header = make(map[string]string)
	}
	s.Header[key] = value
	return s
}

func (s *HttpClient) setRequestHeader(r *http.Request) *http.Request {
	if r == nil || s.Header == nil || len(s.Header) == 0 {
		return r
	}
	for k, v := range s.Header {
		r.Header.Set(k, v)
	}
	return r
}

func (s *HttpClient) Post(uri string, data []byte, header map[string][]string) (*http.Response, error) {
	var body io.Reader = nil
	if len(data) != 0 {
		body = strings.NewReader(string(data))
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	if len(header) != 0 {
		for k, vs := range header {
			for _, v := range vs {
				s.setRequestHeader(req).Header.Set(k, v)
			}
		}
	}
	return s.client.Do(req)
}

func (s *HttpClient) PostJSON(uri string, data interface{}) (*http.Response, error) {
	var body io.Reader = nil
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(b))
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	s.setRequestHeader(req).Header.Set("Content-Type", "application/json")
	return s.client.Do(req)
}

func (s *HttpClient) PostForm(uri string, data map[string]string) (*http.Response, error) {
	value := make(url.Values)
	for k, v := range data {
		value.Set(k, v)
	}
	return http.PostForm(uri, value)
}

func (s *HttpClient) Put(uri string, data interface{}) (*http.Response, error) {
	var body io.Reader = nil
	if data != nil {
		j, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(j))
	}
	req, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return nil, err
	}
	s.setRequestHeader(req).Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return s.client.Do(req)
}

func (s *HttpClient) Get(uri string) (*http.Response, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(s.setRequestHeader(req))
}

func (s *HttpClient) Delete(uri string, data interface{}) (*http.Response, error) {
	var body io.Reader = nil
	if data != nil {
		j, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(j))
	}
	req, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return nil, err
	}
	s.setRequestHeader(req).Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return s.client.Do(req)
}

func (s *HttpClient) PostFile(uri string, filepath string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fh, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	fileWriter, err := bodyWriter.CreateFormFile("file", path.Base(fh.Name()))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, bodyBuf)
	if err != nil {
		return nil, err
	}
	s.setRequestHeader(req).Header.Set("Content-Type", contentType)
	return s.client.Do(req)
}

func (s *HttpClient) GetFile(uri string, filepath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	s.setRequestHeader(req).Header.Set("Content-Type", "multipart/form-data")
	resp, err := s.client.Do(req)
	defer resp.Body.Close()

	err = os.MkdirAll(path.Dir(filepath), 0766)
	if err != nil {
		return resp, err
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return resp, err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return resp, err
}
