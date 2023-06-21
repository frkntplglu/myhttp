package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

func New(timeout time.Duration) *HttpClient {
	client := &http.Client{
		Timeout: timeout,
	}

	return &HttpClient{
		Client: client,
	}
}

func (c *HttpClient) Get(url string) ([]byte, error) {
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request failed with status code %d", resp.StatusCode)

	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
