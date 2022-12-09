package ggpt3

import (
	"context"
	"net/http"
)

func (c *Client) RequestImages(ctx context.Context, ir *ImageRequest) (*ImageResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, nil, "/images/generations", ir)
	if err != nil {
		return nil, err
	}

	var iResp ImageResponse
	err = c.sendRequest(req, &iResp)
	return &iResp, err
}

func (c *Client) RequestImageEdits(ctx context.Context, ir *ImageRequest) (*ImageResponse, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "multipart/form-data")

	req, err := c.newRequest(ctx, http.MethodPost, headers, "/images/edits", ir)
	if err != nil {
		return nil, err
	}

	var iResp ImageResponse
	err = c.sendRequest(req, &iResp)
	return &iResp, err
}

func (c *Client) RequestImageVariations(ctx context.Context, ir *ImageRequest) (*ImageResponse, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "multipart/form-data")
	req, err := c.newRequest(ctx, http.MethodPost, headers, "/images/variations", ir)
	if err != nil {
		return nil, err
	}

	var iResp ImageResponse
	err = c.sendRequest(req, &iResp)
	return &iResp, err
}
