package ggpt3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiBaseUrlv1 = "https://api.openai.com/v1"
)

type Client struct {
	baseUrl    string
	apiKey     string
	headers    http.Header
	httpClient *http.Client
	orgId      string
}

func NewClient(apiKey string) *Client {
	headers := make(http.Header)
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	headers.Set("Content-Type", "application/json")

	return &Client{
		baseUrl: apiBaseUrlv1,
		apiKey:  apiKey,
		headers: headers,
		httpClient: &http.Client{
			Timeout:   time.Minute,
			Transport: &http.Transport{},
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf8")
	}

	if c.orgId != "" {
		req.Header.Set("OpenAI-Organization", c.orgId)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var apiErr ErrorResponse
		err := json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil || apiErr.Error == nil {
			return fmt.Errorf("error, status code %d", res.StatusCode)
		}
		return fmt.Errorf("error, status code %d, message: %s", res.StatusCode, apiErr.Error.Message)
	}

	if v != nil {
		if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) fullRequestUrl(prefix string) string {
	return fmt.Sprintf("%s%s", c.baseUrl, prefix)
}

func (c *Client) newRequest(ctx context.Context, httpMethod string, urlPrefix string, body interface{}) (*http.Request, error) {

	req, err := http.NewRequest(httpMethod, c.fullRequestUrl(urlPrefix), nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	req = req.WithContext(ctx)
	return req, nil
}
