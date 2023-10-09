package types

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GithubClient struct {
	HttpClient *http.Client
}

func NewGithubClientClient() *GithubClient {
	return &GithubClient{
		HttpClient: &http.Client{},
	}
}

type HttpInterface interface {
	Get(url string) (resp *http.Response, err error)
	Post(url string, payload interface{}) (resp *http.Response, err error)
	Patch(url string, payload interface{}) (resp *http.Response, err error)
	Put(url string, payload interface{}) (resp *http.Response, err error)
	Delete(url string) (resp *http.Response, err error)
}

func (c *GithubClient) Get(url string) (resp *http.Response, err error) {
	return c.do(http.MethodGet, url, nil)
}

func (c *GithubClient) Post(url string, payload interface{}) (resp *http.Response, err error) {
	return c.do(http.MethodPost, url, payload)
}

func (c *GithubClient) Patch(url string, payload interface{}) (resp *http.Response, err error) {
	return c.do(http.MethodPatch, url, payload)
}

func (c *GithubClient) Delete(url string) (resp *http.Response, err error) {
	return c.do(http.MethodDelete, url, nil)
}

func (c *GithubClient) Put(url string, payload interface{}) (resp *http.Response, err error) {
	return c.do(http.MethodPut, url, payload)
}

func (c *GithubClient) do(method string, url string, payload interface{}) (resp *http.Response, err error) {
	request := &http.Request{}
	if method == http.MethodGet || method == http.MethodDelete {
		request, err = http.NewRequest(method, url, nil)
	} else {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest(method, url, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
	}
	GetCommonHeader(request)
	return c.HttpClient.Do(request)
}
