package ggpt3

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) RequestImages(ctx context.Context, ir *ImageRequest) (*ImageResponse, error) {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json; charset=utf-8")

	req, err := c.newRequest(ctx, http.MethodPost, headers, "/images/generations", ir)
	if err != nil {
		return nil, err
	}

	var iResp ImageResponse
	err = c.sendRequest(req, &iResp)
	return &iResp, err
}

func (c *Client) RequestImageEdits(ctx context.Context, ir *ImageRequest) (*ImageResponse, error) {
	if ir.Image == "" {
		return nil, fmt.Errorf("image is a required field for an edits request")
	}
	if ir.Prompt == "" {
		return nil, fmt.Errorf("prompt is a required field for an edits request")
	}

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
	if ir.Image == "" {
		return nil, fmt.Errorf("image is a required field for a variation request")
	}

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
