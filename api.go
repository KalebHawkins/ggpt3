package ggpt3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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

func (c *Client) newRequest(ctx context.Context, httpMethod string, httpHeaders http.Header, urlPrefix string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, c.fullRequestUrl(urlPrefix), nil)
	if err != nil {
		return nil, err
	}

	if httpHeaders != nil {
		req.Header = httpHeaders
	}

	if body != nil {
		contentType := req.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			reqBody, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		case "multipart/form-data":
			fd, ct, err := c.newMultiPartForm(body.(*ImageRequest))
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", ct)
			req.Body = io.NopCloser(bytes.NewBuffer(fd))
		}
	}

	req = req.WithContext(ctx)
	return req, nil
}

func (c *Client) newMultiPartForm(ir *ImageRequest) ([]byte, string, error) {
	f, err := os.Open(ir.Image)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	fw, err := mw.CreateFormFile("image", ir.Image)
	if err != nil {
		return nil, "", err
	}

	_, err = io.Copy(fw, f)
	if err != nil {
		return nil, "", err
	}

	if ir.Mask != "" {
		f, err := os.Open(ir.Mask)
		if err != nil {
			return nil, "", err
		}
		fw, err := mw.CreateFormFile("mask", ir.Mask)
		if err != nil {
			return nil, "", err
		}
		_, err = io.Copy(fw, f)
		if err != nil {
			return nil, "", err
		}
	}
	if ir.Prompt != "" {
		mw.WriteField("prompt", ir.Prompt)
	}
	if ir.N > 0 {
		mw.WriteField("n", fmt.Sprintf("%d", ir.N))
	}
	if ir.Size != "" {
		mw.WriteField("size", ir.Size)
	}
	if ir.ResponseFormat != "" {
		mw.WriteField("response_format", ir.ResponseFormat)
	}
	if ir.User != "" {
		mw.WriteField("user", ir.User)
	}
	mw.Close()

	return body.Bytes(), mw.FormDataContentType(), err
}
