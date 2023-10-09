package types

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CloudClubClient struct {
	HttpClient *http.Client
}

type HttpInterface interface {
	Get(url string) (resp *http.Response, err error)
	Post(url string, payload interface{}) (resp *http.Response, err error)
	Patch(url string, payload interface{}) (resp *http.Response, err error)
	Put(url string, payload interface{}) (resp *http.Response, err error)
	Delete(url string) (resp *http.Response, err error)
}

func NewCloudClubClient() HttpInterface {
	c := &CloudClubClient{}
	return c
}

func (c *CloudClubClient) Get(url string) (resp *http.Response, err error) {
	return c.do(http.MethodGet, url, nil)
}

func (c *CloudClubClient) Post(url string, payload interface{}) (resp *http.Response, err error) {
	return c.do(http.MethodPost, url, payload)
}

func (c *CloudClubClient) Patch(url string, payload interface{}) (resp *http.Response, err error) {
	return c.do(http.MethodPatch, url, payload)
}

func (c *CloudClubClient) Delete(url string) (resp *http.Response, err error) {
	return c.do(http.MethodDelete, url, nil)
}

func (c *CloudClubClient) Put(url string, payload interface{}) (resp *http.Response, err error) {
	return c.do(http.MethodPut, url, nil)
}

func (c *CloudClubClient) do(method string, url string, payload interface{}) (resp *http.Response, err error) {
	var request *http.Request
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

	return c.HttpClient.Do(request)
}
