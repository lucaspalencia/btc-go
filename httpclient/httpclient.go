package httpclient

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	url string
}

func (httpClient *HttpClient) Get() ([]byte, error) {
	res, err := http.Get(httpClient.url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func New(url string) HttpClient {
	client := HttpClient{url: url}

	return client
}
