package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Client struct {
	httpClient     http.Client
	baseAddress    string
	defaultHeaders map[string]string
}

func NewClient(baseAddress string, defaultHeaders map[string]string) *Client {
	logrus.Debugf("Creating new client with base address: %s", baseAddress)
	if baseAddress[len(baseAddress)-1] != '/' {
		baseAddress += "/"
	}

	logrus.Debugf("Base address: %s", baseAddress)

	if _, ok := defaultHeaders["Content-Type"]; !ok {
		defaultHeaders["Content-Type"] = "application/json"
	}

	logrus.Debugf("Default headers: %+v", defaultHeaders)

	return &Client{
		httpClient:     http.Client{},
		baseAddress:    baseAddress,
		defaultHeaders: defaultHeaders,
	}
}

func (c *Client) AddHeader(key, value string) {
	c.defaultHeaders[key] = value
}

func (c *Client) RemoveHeader(key string) {
	delete(c.defaultHeaders, key)
}

func (c *Client) GET(path string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.baseAddress, path), nil)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("GET request: %+v", req)

	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	return c.httpClient.Do(req)
}

func (c *Client) POST(path string, headers map[string]string, body interface{}) (*http.Response, error) {
	var reqBodyBytes []byte
	if body != nil {
		var err error
		reqBodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.baseAddress, path), bytes.NewReader(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return c.httpClient.Do(req)
}

func (c *Client) PUT(path string, headers map[string]string, body interface{}) (*http.Response, error) {
	var reqBodyBytes []byte
	if body != nil {
		var err error
		reqBodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.baseAddress, path), bytes.NewReader(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return c.httpClient.Do(req)
}

func (c *Client) PATCH(path string, headers map[string]string, body interface{}) (*http.Response, error) {
	var reqBodyBytes []byte
	if body != nil {
		var err error
		reqBodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s%s", c.baseAddress, path), bytes.NewReader(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return c.httpClient.Do(req)
}

func (c *Client) DELETE(path string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", c.baseAddress, path), nil)
	if err != nil {
		return nil, err
	}

	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return c.httpClient.Do(req)
}
